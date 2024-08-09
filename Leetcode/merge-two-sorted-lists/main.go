package main

import "fmt"

type ListNode struct {
	Val  *int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) []int {
	MaxLenght := CheckLenListNode(list1, list2)
	var array []int
	for i := 0; i < MaxLenght; i++ {
		array = append(array, *list1.Val)
		list1 = list1.Next
		array = append(array, *list2.Val)
		list2 = list2.Next
	}

	return array
}

func mergeTwoListsFinal(list1 *ListNode, list2 *ListNode) *ListNode {
	MaxLenght := CheckLenListNode(list1, list2)

	if list1.Val == nil && list2.Val == nil {
		return &ListNode{}
	}

	var array []int
	for i := 0; i < MaxLenght; i++ {
		array = append(array, *list1.Val)
		list1 = list1.Next
		array = append(array, *list2.Val)
		list2 = list2.Next
	}

	List := &ListNode{Val: &array[0]}
	current := List
	for i := 1; i < len(array); i++ {
		nextNode := &ListNode{Val: &array[i]}
		current.Next = nextNode
		current = nextNode
	}

	return List
}

func main() {
	//ListNode2 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	//ListNode1 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	ListNode2 := ListNode{}
	ListNode1 := ListNode{}

	//fmt.Println(mergeTwoLists(&ListNode1, &ListNode2))
	//fmt.Println(mergeTwoListsFinal(&ListNode1, &ListNode2))

	mergedList := mergeTwoListsFinal(&ListNode1, &ListNode2)
	fmt.Println(show(mergedList))

}

func show(list *ListNode) []int {
	var array []int
	if list.Val == nil {
		return []int{}
	} else {
		for list.Val != nil {
			array = append(array, *list.Val)
			list = list.Next
		}
	}

	return array
}

func CheckLenListNode(list1 *ListNode, list2 *ListNode) int {
	firstLen := show(list1)
	secondLen := show(list2)

	if len(firstLen) > len(secondLen) {
		return len(firstLen)
	}
	if len(firstLen) < len(secondLen) {
		return len(secondLen)
	} else {
		return len(secondLen)
	}
}
