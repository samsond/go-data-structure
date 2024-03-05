package heap

import "testing"

func TestHeap_Insert(t *testing.T) {

	heap := &Heap[int]{data: []int{}, less: func(a int, b int) bool {
		return a < b
	}}

	values := []int{50, 15, 3, 12, 8, 16, 87, 25, 88, 100}

	for _, v := range values {
		heap.Insert(v)
	}

	expectedRoot := 100

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}

}

func TestHeap_Delete(t *testing.T) {
	heap := &Heap[int]{data: []int{}, less: func(a int, b int) bool {
		return a < b
	}}

	values := []int{50, 15, 3, 12, 8, 16, 87, 25, 88, 100}

	for _, v := range values {
		heap.Insert(v)
	}

	expectedRoot := 100

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}

	heap.Delete()
	expectedRoot = 88

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}

	heap.Delete()
	expectedRoot = 87

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}

}

func TestHeap_Insert_Min_Heap(t *testing.T) {
	heap := &Heap[int]{data: []int{}, less: func(a int, b int) bool {
		return a > b
	}}

	values := []int{50, 15, 3, 12, 8, 16, 87, 25, 88, 100}

	for _, v := range values {
		heap.Insert(v)
	}

	expectedRoot := 3

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}

	heap.Delete()
	expectedRoot = 8

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}

	heap.Delete()
	expectedRoot = 12

	if heap.data[0] != expectedRoot {
		t.Errorf("expected root to be %d; got %d", expectedRoot, heap.data[0])
	}
}
