package main

import "fmt"

type SSLNode struct {
	next  *SSLNode
	value int
}

func (sNode *SSLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SSLNode) GetValue() int {
	return sNode.value
}

func NewSSLNode() *SSLNode {
	return new(SSLNode)
}

func main() {
	node := NewSSLNode()
	node.SetValue(4)
	fmt.Println("Value of Node :", node.GetValue())
}
