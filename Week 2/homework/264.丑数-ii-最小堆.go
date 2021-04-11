/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
// 定义小顶堆，通过实现golang heap 接口
var factors = []int{2, 3, 5}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthUglyNumber(n int) int {
	// 定义最小堆
	h := &hp{sort.IntSlice{1}}
	// 维护seen，避免重复
	seen := map[int]struct{}{1: {}}
	//
	for i := 1; ; i++ {
		// 取出堆顶元素 x
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			// 2x,3x,5x放入堆中
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}

}

// @lc code=end

