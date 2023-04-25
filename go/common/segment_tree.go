package common


type SegmentTreeNode struct {
	left, right int
	sum int
	leftChild, rightChild *SegmentTreeNode
}

type SegmentTree struct {
	Root *SegmentTreeNode
}

func (tree *SegmentTree) Build(array []int) {
	tree.Root = build(array, 0, len(array) - 1)
}

func (tree *SegmentTree) Query(left, right int) int {
	return query(tree.Root, left, right)
}

// 只提供更新操作，像插入、删除要改变数组长度大小的操作，请重新调用Build()
// 虽然Update和Build一样的时间复杂度，但是Update不需要创建分配大量的TreeNode空间消耗时间
// 更新数组中某个索引的值
func (tree *SegmentTree) Update(index, value int) {
	update(tree.Root, index, value)
}

func build(array []int, left, right int) *SegmentTreeNode {
	if left == right { // 只有1个元素时回溯，自底向上建立
		return &SegmentTreeNode{left: left, right: right, sum: array[left]}
	}
    
	// 和归并排序有相似地方
	mid := (left + right) / 2  // split from center
	leftChild := build(array, left, mid)
	rightChild := build(array, mid + 1, right)
	return &SegmentTreeNode{
		left: left, 
		right: right, 
		sum: leftChild.sum + leftChild.sum,  // 当前节点的sum等于它左右节点的sum之和
		leftChild: leftChild, 
		rightChild: rightChild,
	}
}

func query(node *SegmentTreeNode, left, right int) int {
	// 到达叶子节点或者left、right和当前节点之间不存在包含关系
	if node == nil || left > node.right || right < node.left {
		return 0
	}
	// 节点完全被left和right包含
	// 为什么取小于等于，不写完全包含？
	// 比如left=3、right=5，左节点为3..4，右节点5...6，
	// 此时左节点的sum要被计算在内，再加上节点5即得结果
	if left <= node.left && right >= node.right {
		return node.sum  // 说明当前节点在计算范围内，返回它的sum
	}
	// 继续往下搜索子节点，因为有可能left和right被分隔在左右节点，所以取和
	return query(node.leftChild, left, right) + query(node.rightChild, left, right)
}

// 若更新的值和原值相同，将白白消耗时间，调用前在原构建数组的对应索引上判断一下
func update(node *SegmentTreeNode, index, newvalue int) {
	if node == nil || node.left > index || node.right < index {
		// 当前节点范围不包含index
		return
	}

	if node.left == node.right && node.left == index {
		// 找到了要更新的叶子节点
		node.sum = newvalue
		return
	}

	// 继续向下搜索
	update(node.leftChild, index, newvalue)
	update(node.rightChild, index, newvalue)
	// 因为节点更新，sum值也要自底向上逐层更新
	node.sum = node.leftChild.sum + node.leftChild.sum
}
