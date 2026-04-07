package ds

type Queue interface {
	Push(v int)
	Pop() int
}

type Heap interface {
	Push(v int)
	Pop() int
}

type BinaryTree interface {
	Push(v int)
	Pop() int
	Search(v int) bool
}

