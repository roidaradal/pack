package ds

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	m := Map[string, int]{
		"apple":  5,
		"orange": 3,
		"banana": 2,
	}
	expString := "{apple: 5, banana: 2, orange: 3}"
	actualString := m.String()
	if actualString != expString {
		t.Errorf("Map.String() = %q, want %q", actualString, expString)
	}
	m2 := m.Copy()
	if reflect.DeepEqual(m, m2) == false {
		t.Errorf("Map.Copy() = %v, want %v", m2, m)
	}
	size := m.Len()
	if size != 3 {
		t.Errorf("Map.Len() = %d, want 3", size)
	}
	notEmpty := m.NotEmpty()
	if notEmpty != true {
		t.Errorf("m.NotEmpty() = %t, want true", notEmpty)
	}

	m.Clear()
	isEmpty := m.IsEmpty()
	if isEmpty != true {
		t.Errorf("m.IsEmpty() = %t, want true", isEmpty)
	}
}
