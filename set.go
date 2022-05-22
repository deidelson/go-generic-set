package set

type SetEntry[V comparable] interface {
	GetKey() V
}

type HashSet[V comparable, T SetEntry[V]] struct {
	entries map[V]T
}

func NewHashSet[V comparable, T SetEntry[V]]() HashSet[V, T] {
	return HashSet[V, T]{
		entries: make(map[V]T),
	}
}

func (s HashSet[V, T]) Add(entry T) {
	s.entries[entry.GetKey()] = entry
}

func (s HashSet[V, T]) IsEmpty() bool {
	return len(s.entries) == 0
}

func (s HashSet[V, T]) Size() int {
	return len(s.entries)
}

//func (s hashSet) Exists[T SetEntry, K any](entry T) bool {
//
//}
