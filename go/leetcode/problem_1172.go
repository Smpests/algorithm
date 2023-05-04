package leetcode


// 题目链接：https://leetcode.cn/problems/dinner-plate-stacks/
// 超时
type DinnerPlates struct {
    stacks [][]int
    capacity int
    length int
    states []int  // 存每个栈的状态，值表示栈顶元素所在索引
}


func Constructor(capacity int) DinnerPlates {
    return DinnerPlates{capacity: capacity}
}


func (this *DinnerPlates) Push(val int)  {
    for idx, _ := range this.states {
        if this.states[idx] < this.capacity - 1{
            this.states[idx]++
            this.stacks[idx][this.states[idx]] = val
            return
        }
    }
    // 上面的循环没有找到未满的stack，则新建一个栈
    // 因val>=1，因此这里可以初始化为0而不影响
    stack := make([]int, this.capacity)
    stack[0] = val
    this.stacks = append(this.stacks, stack)
    this.states = append(this.states, 0)
    this.length += 1
}


func (this *DinnerPlates) Pop() int {
    for i := this.length - 1; i > -1; i-- {
        if this.states[i] > -1 {
            return this.getPopValAt(i)
        }
    }
    return -1
}


// 获取某索引的栈顶元素，index有效性在调用前判断
func (this *DinnerPlates) getPopValAt(index int) int {
    stackIndex := this.states[index]
    // fmt.Println(index)
    // fmt.Println(this.stacks)
    // fmt.Println(this.states)
    popVal := this.stacks[index][stackIndex]
    this.stacks[index][stackIndex] = 0
    this.states[index]--
    return popVal
}


func (this *DinnerPlates) PopAtStack(index int) int {
    if index >= this.length || index < 0 {
        return -1
    }
    if this.states[index] > -1 {
        return this.getPopValAt(index)
    }
    return -1
}


/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */