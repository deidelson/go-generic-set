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

func (s HashSet[V, T]) Put(entry T) {
	s.entries[entry.GetKey()] = entry
}

func (s HashSet[V, T]) Len() int {
	return len(s.entries)
}

//func (s hashSet) Exists[T SetEntry, K any](entry T) bool {
//
//}

type Student struct {
	Dni      string
	Edad     int
	Nombre   string
	Apellido string
}

type StudentKey struct {
	Dni string
}

func NewStudent(dni string) Student {
	return Student{
		Dni:      dni,
		Nombre:   "nombre",
		Apellido: "apellido",
		Edad:     23,
	}
}

func (s Student) GetKey() StudentKey {
	return StudentKey{
		Dni: s.Dni,
	}
}

/* func main() {
	set := NewHashSet[StudentKey, Student]()

	set.Put(NewStudent("1"))
	set.Put(NewStudent("2"))
	set.Put(NewStudent("3"))
	set.Put(NewStudent("1"))

	log.Println(set.Len())

} */
