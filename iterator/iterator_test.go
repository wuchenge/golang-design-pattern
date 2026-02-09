package iterator

import "testing"

func TestIterator(t *testing.T) {
	aggregate := NewNumbers(1, 10)
	IteratorPrint(aggregate.Iterator())
	t.Log("testing...")
}
