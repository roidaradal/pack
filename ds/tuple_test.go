package ds

import "testing"

func TestTuple2(t *testing.T) {
	v1, v2 := "apple", 5
	t1 := NewTuple2(v1, v2)
	if t1.V1 != v1 || t1.V2 != v2 {
		t.Errorf("Tuple2.V1, V2 = %v, %v, want %v, %v", t1.V1, t1.V2, v1, v2)
	}
	a, b := t1.Values()
	if a != v1 || b != v2 {
		t.Errorf("Tuple2.Values() = %v, %v, want %v, %v", a, b, v1, v2)
	}
}
