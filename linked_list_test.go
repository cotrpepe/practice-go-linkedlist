package linked_list

import (
	"testing"
)

func TestSize(t *testing.T) {
	assert := func(actual int, expected int) {
		if actual != expected {
			t.Errorf("Size() = %v, want %v", actual, expected)
		}
	}

	list := new(LinkedList)
	assert(list.Size(), 0)

	list.Add(1)
	assert(list.Size(), 1)

	list.Add(2)
	assert(list.Size(), 2)

	list.Delete(1)
	assert(list.Size(), 1)

	list.Delete(2)
	assert(list.Size(), 0)
}

func TestEmpty(t *testing.T) {
	assert := func(actual bool, expected bool) {
		if actual != expected {
			t.Errorf("Empty() = %v, want %v", actual, expected)
		}
	}

	list := new(LinkedList)
	assert(list.Empty(), true)

	list.Add(1)
	assert(list.Empty(), false)

	list.Delete(1)
	assert(list.Empty(), true)
}

func TestDelete(t *testing.T) {
	err := func(actual error) {
		if actual == nil {
			t.Errorf("Delete() = %v, want errors", actual)
		}
	}
	no_err := func(actual error) {
		if actual != nil {
			t.Errorf("Delete() = %v, want nil", actual)
		}
	}

	list := new(LinkedList)
	err(list.Delete(1))

	list.Add(1)
	list.Add(2)
	err(list.Delete(3))
	no_err(list.Delete(1))
	no_err(list.Delete(2))
	err(list.Delete(2))
}

