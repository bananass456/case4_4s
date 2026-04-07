package ds

import "testing"

type queueStub struct{}

func (queueStub) Push(v int) {}
func (queueStub) Pop() int   { return 0 }

type heapStub struct{}

func (heapStub) Push(v int) {}
func (heapStub) Pop() int   { return 0 }

type treeStub struct{}

func (treeStub) Push(v int)        {}
func (treeStub) Pop() int          { return 0 }
func (treeStub) Search(v int) bool { return false }

func TestQueueSatisfiesInterface(t *testing.T) {
	var _ Queue = queueStub{}
}

func TestHeapSatisfiesInterface(t *testing.T) {
	var _ Heap = heapStub{}
}

func TestBinaryTreeSatisfiesInterface(t *testing.T) {
	var _ BinaryTree = treeStub{}
}

