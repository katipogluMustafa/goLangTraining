package main

import "fmt"

type SSLNode struct {
	next  *SSLNode
	value int
}

// here we didn't give the value to the one we want to show, rather we gave the copy of the node which just defined in this scope.
func (sNode SSLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode SSLNode) GetValue() int {
	return sNode.value
}

func NewSSLNode() SSLNode {
	return SSLNode{}
}

func main() {
	node := NewSSLNode()
	node.SetValue(4)
	fmt.Println("Value of Node :", node.GetValue())
}
