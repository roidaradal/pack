package ds

import (
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	// NewQueue
	q := NewQueue[int]()
	actualString, wantString := q.String(), "[]"
	if actualString != wantString {
		t.Errorf("Queue.String() = %q, want %q", actualString, wantString)
	}
	actualCount, wantCount := q.Len(), 0
	if actualCount != wantCount {
		t.Errorf("Queue.Len() = %d, want %d", actualCount, wantCount)
	}
	actualFlag := q.IsEmpty()
	if actualFlag != true {
		t.Errorf("Queue.IsEmpty() = %t, want %t", actualFlag, true)
	}
	actualFlag = q.NotEmpty()
	if actualFlag != false {
		t.Errorf("Queue.NotEmpty() = %t, want %t", actualFlag, false)
	}
	// NewQueueFrom
	items := []int{1, 2, 3}
	q = NewQueueFrom(items)
	actualString, wantString = q.String(), "[1 2 3]"
	if actualString != wantString {
		t.Errorf("Queue.String() = %q, want %q", actualString, wantString)
	}
	actualCount, wantCount = q.Len(), len(items)
	if actualCount != wantCount {
		t.Errorf("Queue.Len() = %d, want %d", actualCount, wantCount)
	}
	actualFlag = q.IsEmpty()
	if actualFlag != false {
		t.Errorf("Queue.IsEmpty() = %t, want %t", actualFlag, false)
	}
	actualFlag = q.NotEmpty()
	if actualFlag != true {
		t.Errorf("Queue.NotEmpty() = %t, want %t", actualFlag, true)
	}
	actualItems := q.Items()
	if slices.Equal(items, actualItems) == false {
		t.Errorf("Queue.Items() = %v, want %v", actualItems, items)
	}
	// Copy
	q2 := q.Copy()
	wantItems, actualItems := q.Items(), q2.Items()
	if slices.Equal(wantItems, actualItems) == false {
		t.Errorf("Queue.Copy.Items() = %v, want %v", actualItems, wantItems)
	}
	// Clear
	q2.Clear()
	actualFlag = q2.IsEmpty()
	if actualFlag != true {
		t.Errorf("Queue.Clear.IsEmpty() = %t, want %t", actualFlag, true)
	}
	// Check original queue is unchanged
	actualItems = q.Items()
	if slices.Equal(items, actualItems) == false {
		t.Errorf("Queue.Items() = %v, want %v", actualItems, items)
	}
}

func TestQueueMethods(t *testing.T) {
	// TODO: Enqueue
	// TODO: Front, MustFront
	// TODO: Dequeue, MustDequeue

}
