// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import "fmt"

type StringNode struct {
	Value string
	Left  *StringNode
	Right *StringNode
}

type StringBinaryTree struct {
	Root *StringNode
}

func (bt *StringBinaryTree) Insert(value string) {
	node := &StringNode{value, nil, nil}
	if bt.Root == nil {
		bt.Root = node
	} else {
		insertStringNode(bt.Root, node)
	}
}

func insertStringNode(node, newNode *StringNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertStringNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertStringNode(node.Right, newNode)
		}
	}
}

func printStringTree(node *StringNode, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printStringTree(node.Left, level)
		fmt.Printf(format+"%v\n", node.Value)
		printStringTree(node.Right, level)
	}
}

func main() {
	var bt StringBinaryTree
	bt.Insert("8")

	bt.Insert("3")
	bt.Insert("11")

	bt.Insert("1")
	bt.Insert("0")
	bt.Insert("2")

	bt.Insert("5")
	bt.Insert("4")
	bt.Insert("6")

	bt.Insert("9")
	bt.Insert("8")
	bt.Insert("10")

	bt.Insert("13")
	bt.Insert("12")
	bt.Insert("14")

	printStringTree(bt.Root, 0)
}