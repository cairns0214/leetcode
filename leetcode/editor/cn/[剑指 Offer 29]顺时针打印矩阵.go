//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。 
//
// 
//
// 示例 1： 
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
// 
//
// 示例 2： 
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
//
// 限制： 
//
// 
// 0 <= matrix.length <= 100 
// 0 <= matrix[i].length <= 100 
// 
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/ 
// Related Topics 数组 矩阵 模拟 👍 317 👎 0

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return make([]int, 0)
	}
	if len(matrix) == 1 {
		return matrix[0]
	}

	if len(matrix[0]) == 1 {
		res := make([]int, len(matrix))
		for i, arr := range matrix {
			res[i] = arr[0]
		}
		return res
	}

	res := make([]int, len(matrix)*len(matrix[0]))
	i := 0
	a, b, c, d := 0, len(matrix[0]) - 1, len(matrix) - 1, 0

	for a < c && d < b {
		for j := d; j < b; j++  {
			res[i] = matrix[a][j]
			i++
		}
		for j := a; j < c; j++  {
			res[i] = matrix[j][b]
			i++
		}
		for j := b; j > d; j--  {
			res[i] = matrix[c][j]
			i++
		}
		for j := c; j > a; j--  {
			res[i] = matrix[j][d]
			i++
		}
		a++
		b--
		c--
		d++
	}
	if a == c {
		for j := d; j <=b; j++ {
			res[i] = matrix[a][j]
			i++
		}
		return res
	}

	if b == d {
		for j := a; j <= c; j++ {
			res[i] = matrix[j][b]
			i++
		}
		return res
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

// [0][0] - [0][n-2]
// [0][n-1] - [m-2][n-1]
// [m-1][n-1] - [m-1][1]
// [m-1][0] - [1][0]
