package dict

import (
	"reflect"
	"testing"
)

func TestDict(t *testing.T) {
	m := map[string]int{
		"apple":  5,
		"orange": 3,
		"banana": 2,
	}
	m2 := Copy(m)
	if reflect.DeepEqual(m, m2) == false {
		t.Errorf("Copy() = %v, want %v", m2, m)
	}
	size := Len(m)
	if size != 3 {
		t.Errorf("Len() = %d, want 3", size)
	}
	notEmpty := NotEmpty(m)
	if notEmpty != true {
		t.Errorf("NotEmpty() = %t, want true", notEmpty)
	}

	Clear(m)
	isEmpty := IsEmpty(m)
	if isEmpty != true {
		t.Errorf("IsEmpty() = %t, want true", isEmpty)
	}
}
