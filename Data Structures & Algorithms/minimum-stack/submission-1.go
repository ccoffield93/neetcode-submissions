type MinStack struct {
	stack []int
    min   []int
}

func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        min:   []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    
    // Determine the current minimum at this stack depth
    n := len(this.min)
    if n == 0 {
        this.min = append(this.min, val)
    } else {
        currentMin := this.min[n-1]
        this.min = append(this.min, min(val, currentMin))
    }
}

func (this *MinStack) Pop() {
    n := len(this.stack)
    if n == 0 {
        return
    }
    this.stack = this.stack[:n-1]
    this.min = this.min[:n-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1]
}
