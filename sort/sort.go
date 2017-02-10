package main

import (
	"fmt"
)

func InsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		key := s[i] //要插入的值
		j := i - 1
		for ; j >= 0 && key < s[j]; j-- {
			s[j+1] = s[j]
		}
		s[j+1] = key
	}
}

func QuickSort(s []int) {
	if len(s) <= 1 {
		return
	}
	low, high := 0, len(s)-1
	key := s[0]
	for low < high {
		for ; low < high && s[high] > key; high-- {
		}
		s[low] = s[high]
		for ; low < high && s[low] < key; low++ {
		}
		s[high] = s[low]
	}
	s[low] = key
	QuickSort(s[:low])
	QuickSort(s[low+1:])
}

func Merge(a, b []int) (result []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return
}

func MergeSort(s []int) []int {
	length := len(s)
	if length <= 1 {
		return s
	}
	mid := length / 2
	left := MergeSort(s[:mid])
	right := MergeSort(s[mid:])
	return Merge(left, right)
}

type tree struct {
	value       int
	left, right *tree
}

func TreeSort(s []int) []int {
	var root *tree
	for _, v := range s {
		root = add(root, v)
	}
	return Inorder(s[:0], root)
}

func Inorder(value []int, root *tree) []int {
	if root != nil {
		value = Inorder(value, root.left)
		value = append(value, root.value)
		value = Inorder(value, root.right)
	}
	return value
}

func add(root *tree, value int) *tree {
	if root == nil {
		root := new(tree)
		root.value = value
		return root
	}
	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func main() {
	s := []int{77, 7, 9, 56, 8, 32, 1, 5, 2}
	//	fmt.Println(s)
	// QuickSort(s)
	//s = MergeSort(s)
	s = TreeSort(s)
	fmt.Println(s)
}
