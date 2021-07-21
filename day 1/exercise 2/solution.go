package main

import "fmt"

type Node struct{
	val string
	left *Node
	right *Node
}

func isOperator(c string) bool {
	if c=="+" || c=="-" || c=="*" || c=="/" || c=="^" {
		return true
	}
	return false
}
func preorder(tree *Node) {
	fmt.Println(tree.val)
	preorder(tree.left)
	preorder(tree.right)
}
func postorder(tree *Node) {
	postorder(tree.left)
	postorder(tree.left)
	fmt.Println(tree.val)
}

func newNode(s string) *Node {
	temp:= new(Node)
	temp.left=nil
	temp.right=nil
	temp.val=s

	return temp
}
func constructTree(str string) *Node {
	var stack[] *Node
	var t,t1,t2 *Node

	for i:=0;i<len(str);i++ {
		if !isOperator(string(str[i])) {
			t=newNode(string(str[i]))
			stack=append(stack,t)
		} else {
			t=newNode(string(str[i]))
			if len(stack)>0 {
				n := len(stack) - 1
				t1 = stack[n]
				stack=stack[:n]
			}
			if len(stack)>0 {
				n := len(stack) - 1
				t2 = stack[n]
				stack = stack[:n]
			}
			t.right=t1
			t.left=t2
			stack=append(stack,t)
		}
	}
	if len(stack)>0 {
		n := len(stack) - 1
		t = stack[n]
		stack = stack[:n]
	}
	return t
}
func main() {
	exp:="a+b-c"
	res:= constructTree(exp)
	fmt.Print("Preorder Traversal : ")
	preorder(res)
	fmt.Println()
	fmt.Print("Postorder Traversal : ")
	postorder(res)
}


