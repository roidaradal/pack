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

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	front := q.MustFront()
	if front != 1 {
		t.Errorf("Enquque.MustFront() = %d, want 1", front)
	}
	q.Enqueue(2)
	q.Enqueue(3)
	front = q.MustFront()
	if front != 1 {
		t.Errorf("Enquque.MustFront() = %d, want 1", front)
	}
	want := []int{1, 2, 3}
	actual := q.Items()
	if slices.Equal(want, actual) == false {
		t.Errorf("Queue.Items() = %v, want %v", actual, want)
	}
}

func TestQueueFront(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Queue.MustFront() did not panic")
		}
	}()
	q := NewQueue[int]()
	front := q.Front()
	if !front.IsNil() {
		t.Errorf("Queue.Front() = %v, want nil", front)
	}
	q.Enqueue(1)
	front = q.Front()
	if front.IsNil() || front.Value() != 1 {
		t.Errorf("Queue.Front() = %v, want 1", front)
	}
	frontItem := q.MustFront()
	if frontItem != 1 {
		t.Errorf("Queue.MustFront() = %d, want 1", frontItem)
	}
	q.Dequeue()
	q.MustFront() // should panic
}

func TestQueueDequeue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Queue.MustDequeue() did not panic")
		}
	}()
	q := NewQueueFrom([]int{1, 2})
	front := q.Dequeue()
	if front.IsNil() || front.Value() != 1 {
		t.Errorf("Queue.Dequeue() = %v, want 1", front)
	}
	frontItem := q.MustDequeue()
	if frontItem != 2 {
		t.Errorf("Queue.MustDequeue() = %d, want 2", frontItem)
	}
	front = q.Dequeue()
	if !front.IsNil() {
		t.Errorf("Queue.Dequeue() = %v, want nil", front)
	}
	q.MustDequeue() // should panic
}
