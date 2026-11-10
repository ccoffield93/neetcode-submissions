func evalRPN(tokens []string) int {
    var stack []int

    for _, token := range tokens {
        switch token {
        case "+":
            b := pop(&stack) // Note: order matters for - and /
            a := pop(&stack)
            push(&stack, a+b)
        case "-":
            b := pop(&stack)
            a := pop(&stack)
            push(&stack, a-b)
        case "*":
            b := pop(&stack)
            a := pop(&stack)
            push(&stack, a*b)
        case "/":
            b := pop(&stack)
            a := pop(&stack)
            push(&stack, a/b)
        default:
            num, _ := strconv.Atoi(token)
            push(&stack, num)
        }
    }
    return pop(&stack)
}

func pop(stack *[]int) int {
    n := len(*stack)
    ret := (*stack)[n-1]
    *stack = (*stack)[:n-1] 
    return ret
}

func push(stack *[]int, toAdd int) {
    *stack = append(*stack, toAdd) 
}