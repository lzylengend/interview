package algo

type MinStack struct {
	Stack    []int
	MinIndex int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 0),
		MinIndex: -1,
	}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)

	if this.MinIndex == -1 || x < this.Stack[this.MinIndex] {
		this.MinIndex = len(this.Stack) - 1
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}

	this.Stack = this.Stack[:len(this.Stack)-1]

	if this.MinIndex == len(this.Stack) {
		this.MinIndex = 0
		for i, v := range this.Stack {
			if v < this.Stack[this.MinIndex] {
				this.MinIndex = i
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}

	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	if this.MinIndex == -1 {
		return 0
	}

	return this.Stack[this.MinIndex]
}
