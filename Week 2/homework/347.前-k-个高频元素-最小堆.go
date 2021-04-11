type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1] // 比较频率
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// 1. 最小堆
	// 统计数字出现次数 数组->频率
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	h := &IntHeap{}
	heap.Init(h)
	for key, val := range hashMap {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).([2]int)[0])
	}
	return result
}