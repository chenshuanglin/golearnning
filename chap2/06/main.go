package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(v, root)
	}
	return appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(value int, root *tree) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}

	if value < root.value {
		root.left = add(value, root.left)
	} else {
		root.right = add(value, root.right)
	}
	return root
}

func modifyTreeValue(t *tree) {
	t.value++
}

func main() {
	arr := []int{5, 2, 5, 78, 3, 4, 5, 9, 0, 1, 5, 3, 7, 4, 4}
	fmt.Println(Sort(arr))

	var myTree *tree
	myTree = new(tree)
	myTree.value = 2
	modifyTreeValue(myTree)
	fmt.Println(myTree.value)
}
