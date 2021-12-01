//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。 
//
// 
//
// 示例： 
//
// 
//输入：nums = [1,2,3,4]
//输出：[1,3,2,4] 
//注：[3,1,2,4] 也是正确的答案之一。 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 50000 
// 0 <= nums[i] <= 10000 
// 
// Related Topics 数组 双指针 排序 👍 175 👎 0

package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func exchange(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i,j := 0, len(nums)-1; i < j ; {
		for nums[i] % 2 == 1 && i < j {
			i++
		}
		for nums[j] % 2 == 0 && i < j{
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
//leetcode submit region end(Prohibit modification and deletion)
