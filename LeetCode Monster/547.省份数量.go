/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */
type UnionFind struct {
	father    map[int]int // 记录父节点
	circleNum int
}

func NewUnionFind() UnionFind {
	return UnionFind{
		father: make(make[int]int)
		circleNum:0
	}
}

func (this *UnionFind) Add(int x) {
	if _,exist := this.father[x];!exist {
		this.father[x] = -1
		this.circleNum++
	}
}

func (this *UnionFind) Find(int x) int {
	root := x
	for father[x] != -1 {
		root = father[root]
	}

	for x != root {
		original_father := father[x]
		father[x] = root
		x = original_father
	}

	return root

}

func (this *UnionFind) IsConnected（x,y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Merge(x,y int) {
	xRoot , yRoot := this.Find(x), this.Find(y)
	if xRoot != yRoot {
		father[x] = yRoot
		this.circleNum--
	}
}

func (this *UnionFind) GetCircleNum() int {
	return this.circleNum
}

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	uf := NewUnionFind()
	for i:= 0; i < len(isConnected); i++{
		uf.Add(i)
		for j:=0;j<len(isConnected);j++{
			if isConnectedp[i][j] {
				uf.Merge(i,j)
			}
		}
	}
	return uf.GetCircleNum
}

// @lc code=end

