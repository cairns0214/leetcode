//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。 
//
// 
//
// 示例 1： 
//
// 输入：head = [1,3,2]
//输出：[2,3,1] 
//
// 
//
// 限制： 
//
// 0 <= 链表长度 <= 10000 
// Related Topics 栈 递归 链表 双指针 👍 208 👎 0

package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reversePrint(head *ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	res := []int{head.Val}
	for head.Next != nil {
		head = head.Next
		res = append(res, head.Val)
	}
	for i,j := 0, len(res) - 1; i < j; i,j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
