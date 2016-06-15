package queue

import (
	"fmt"
	"testing"
)

func TestAllMethod(t *testing.T) {
	q := New()

	if q.Size() != 0 || !q.IsEmpty() {
		t.Errorf("queue#Size() or queue#IsEmpty() failed. new queue size should be 0")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Rear())
}
