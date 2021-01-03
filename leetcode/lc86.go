package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	realHead := new(ListNode)
	realHead.Next = head
	// 找到第一个大于或等于x节点的位置
	beforeInsert, insert := realHead, head
	for insert != nil && insert.Val < x {
		beforeInsert = insert
		insert = insert.Next
	}
	// 没有找到大于或等于x的节点
	if insert == nil {
		return head
	}
	// 继续遍历链表，把小于x的节点插到第一个大于等于x节点的位置前面
	beforeP, p := insert, insert.Next
	for p != nil {
		if p.Val < x {
			beforeP.Next = p.Next
			p.Next = insert
			beforeInsert.Next = p
			beforeInsert = p
			p = beforeP.Next
		} else {
			beforeP = p
			p = p.Next
		}
	}
	return realHead.Next
}
