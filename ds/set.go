package ds

type Set[T comparable] struct {
	items Map[T, struct{}]
}

// Len returns the Set size
func (s Set[T]) Len() Int {
	return s.items.Len()
}

// IsEmpty checks if the Set is empty
func (s Set[T]) IsEmpty() Boolean {
	return s.items.IsEmpty()
}

// NotEmpty checks if the Set is not empty
func (s Set[T]) NotEmpty() Boolean {
	return s.items.NotEmpty()
}

// Copy creates a new Set with copied items
func (s Set[T]) Copy() Set[T] {
	return Set[T]{items: s.items.Copy()}
}

// Items returns the Set items, in arbitrary order
func (s Set[T]) Items() List[T] {
	return s.items.Keys()
}
