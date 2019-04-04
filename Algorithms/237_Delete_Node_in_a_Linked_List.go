// 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
// 
// 说明:
// 链表至少包含两个节点。
// 链表中所有节点的值都是唯一的。
// 给定的节点为非末尾节点并且一定是链表中的一个有效节点。
// 不要从你的函数中返回任何结果。

package main

func test237() {

}

// 把 node下一节点复制到node，把下一节点跳过！
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}