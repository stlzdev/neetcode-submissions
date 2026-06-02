type MinStack struct {
	stack []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack : []int{},
		minstack : []int{},
	}
}

func (this *MinStack) Push(val int) {
	clen := len(this.minstack)
	if len(this.minstack) == 0 {
		this.minstack = append(this.minstack, val)
	} else {
		currmin := min(this.minstack[clen - 1], val)
		this.minstack = append(this.minstack, currmin)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	clen := len(this.stack)
	this.stack = this.stack[:clen - 1]
	this.minstack = this.minstack[:clen - 1]
}

func (this *MinStack) Top() int {
	clen := len(this.stack)
	return this.stack[clen - 1]
}

func (this *MinStack) GetMin() int {
	clen := len(this.stack)
	return this.minstack[clen - 1]
}