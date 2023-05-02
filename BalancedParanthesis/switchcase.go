func BalancedBrackets(s string) bool {
    var stack []byte
    if len(s) == 0 {
        return false
    }
    l := len(s)
    for i := 0; i < l; i++ {
        switch s[i] {
        case '{', '(', '[':
            stack = append(stack, s[i])
        case '}':
            if len(stack) == 0 || stack[len(stack)-1] != '{' {
                return false
            }
            stack = stack[:len(stack)-1]
        case ')':
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
        case ']':
            if len(stack) == 0 || stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
