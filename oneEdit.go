package main



func OneEdit(stringOne string, stringTwo string) bool {
    edit := 0
    if (len(stringOne)-len(stringTwo)) > 1 {
        return false
    }
    if len(stringOne) == len(stringTwo) {
        for i, j := 0,0; i < len(stringOne) && j < len(stringTwo); i, j = i+1, j+1 {
            if stringOne[i] != stringTwo[j] {
                edit++
            }
            if edit > 1 {
                return false
            }
        }
    } else {
        for i, j := 0, 0; i < len(stringOne) && j < len(stringTwo); {
            if stringOne[i] != stringTwo[j] {
                if len(stringOne) < len(stringTwo) {
                    edit++
                    j++
                } else {
                    edit++
                    i++
                }
            } else {
                i++
                j++
            }
            if edit > 1 {
                return false
            }
        }
    }
    return true
}
