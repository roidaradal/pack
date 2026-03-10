package ds

// Tuple2 is a tuple with two types
type Tuple2[A, B any] struct {
	V1 A
	V2 B
}

// NewTuple2 creates a new Tuple2
func NewTuple2[A, B any](v1 A, v2 B) Tuple2[A, B] {
	return Tuple2[A, B]{v1, v2}
}

// Values returns the unpacked Tuple2 values
func (t Tuple2[A, B]) Values() (A, B) {
	return t.V1, t.V2
}
