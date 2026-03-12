package ds

import (
	"slices"
	"testing"
)

func TestStack(t *testing.T) {
	// NewStack
	s := NewStack[int]()
	actualString, wantString := s.String(), "[]"
	if actualString != wantString {
		t.Errorf("Stack.String() = %q, want %q", actualString, wantString)
	}
	actualCount, wantCount := s.Len(), 0
	if actualCount != wantCount {
		t.Errorf("Stack.Len() = %d, want %d", actualCount, wantCount)
	}
	actualFlag := s.IsEmpty()
	if actualFlag != true {
		t.Errorf("Stack.IsEmpty() = %t, want %t", actualFlag, true)
	}
	actualFlag = s.NotEmpty()
	if actualFlag != false {
		t.Errorf("Stack.NotEmpty() = %t, want %t", actualFlag, false)
	}
	// NewStackFrom
	items := []int{1, 2, 3}
	s = NewStackFrom[int](items)
	actualString, wantString = s.String(), "[1 2 3]"
	if actualString != wantString {
		t.Errorf("Stack.String() = %q, want %q", actualString, wantString)
	}
	actualCount, wantCount = s.Len(), len(items)
	if actualCount != wantCount {
		t.Errorf("Stack.Len() = %d, want %d", actualCount, wantCount)
	}
	actualFlag = s.IsEmpty()
	if actualFlag != false {
		t.Errorf("Stack.IsEmpty() = %t, want %t", actualFlag, false)
	}
	actualFlag = s.NotEmpty()
	if actualFlag != true {
		t.Errorf("Stack.NotEmpty() = %t, want %t", actualFlag, true)
	}
	actualItems := s.Items()
	if slices.Equal(items, actualItems) == false {
		t.Errorf("Stack.Items() = %v, want %v", actualItems, items)
	}
	// Copy
	s2 := s.Copy()
	wantItems, actualItems := s.Items(), s2.Items()
	if slices.Equal(wantItems, actualItems) == false {
		t.Errorf("Stack.Copy.Items() = %v, want %v", actualItems, wantItems)
	}
	// Clear
	s2.Clear()
	actualFlag = s2.IsEmpty()
	if actualFlag != true {
		t.Errorf("Stack.Clear.IsEmpty() = %t, want %t", actualFlag, true)
	}
	// Check original stack is unchanged
	actualItems = s.Items()
	if slices.Equal(items, actualItems) == false {
		t.Errorf("Stack.Items() = %v, want %v", actualItems, items)
	}
}

func TestStackTop(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Stack.MustTop() did not panic")
		}
	}()
	s := NewStack[int]()
	top := s.Top()
	if !top.IsNil() {
		t.Errorf("Stack.Top() = %v, want nil", top)
	}
	s.Push(1)
	top = s.Top()
	if top.IsNil() || top.Value() != 1 {
		t.Errorf("Stack.Top() = %v, want 1", top)
	}
	topItem := s.MustTop()
	if topItem != 1 {
		t.Errorf("Stack.MustTop() = %v, want 1", topItem)
	}
	s.Pop()
	s.MustTop() // should panic
}

func TestStackPop(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Stack.MustPop() did not panic")
		}
	}()
	s := NewStackFrom[int]([]int{1, 2})
	top := s.Pop()
	if top.IsNil() || top.Value() != 2 {
		t.Errorf("Stack.Pop() = %v, want 2", top)
	}
	topItem := s.MustPop()
	if topItem != 1 {
		t.Errorf("Stack.MustPop() = %v, want 1", topItem)
	}
	top = s.Pop()
	if !top.IsNil() {
		t.Errorf("Stack.Pop() = %v, want nil", top)
	}
	s.MustPop() // should panic
}
