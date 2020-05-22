package helper

type Node struct {
	Prev *Node
	Next *Node
	Cur  interface{}
}
