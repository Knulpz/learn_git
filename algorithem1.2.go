package main

//Leetode21 合并两个有序链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.使用哨兵节点
	cur := &ListNode{} //使用cur记录合成后的链表进度
	sentinel := cur    //sentinel和cur指向同一个初始节点

	//2.遍历两个链表
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}

	}

	return sentinel.Next
}
