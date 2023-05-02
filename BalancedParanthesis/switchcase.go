func BalancedBrackets(s string) bool {
    var arr []string
	if len(s) == 0 {
        return false
    }
    for _, c := range s {
        val := string(c)
        switch val {
        case "{", "(", "[":
            arr = append(arr, val)
        case "}":
            if len(arr) == 0 || arr[len(arr)-1] != "{" {
                return false
            }
            arr = arr[:len(arr)-1]
        case ")":
            if len(arr) == 0 || arr[len(arr)-1] != "(" {
                return false
            }
            arr = arr[:len(arr)-1]
        case "]":
            if len(arr) == 0 || arr[len(arr)-1] != "[" {
                return false
            }
            arr = arr[:len(arr)-1]
        }
    }
    return len(arr) == 0
}
