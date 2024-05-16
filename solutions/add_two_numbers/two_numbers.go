package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		remainder := sum % 10

		current.Next = &ListNode{Val: remainder}
		current = current.Next
	}

	return dummy.Next
}

func main() {
	// Test case 1: 342 + 465 = 807
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)
	fmt.Print("Result of test case 1: ")
	printList(result)
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
	fmt.Println("nil")
}
