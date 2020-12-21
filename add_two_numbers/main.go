package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head

	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0

		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}

		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}

		num := n1 + n2 + carry
		carry = num / 10

		// ループの初回のcur.Nextはhead.Next
		// ループの2回目cur.Nextはhead.Next.Next
		cur.Next = &ListNode{Val: num % 10, Next: nil}
		// ここでcur自体は初期化
		cur = cur.Next
	}

	return head.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	output := addTwoNumbers(l1, l2)

	tmp := output
	for tmp != nil {
		tmp = tmp.Next
	}

}
