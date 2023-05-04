package main

import "strconv"

func RunLengthEncoding(str string) string {
    count := 1
    result := ""
    for i :=0; i < len(str);i++ {
        for i+1 < len(str) && str[i] == str[i+1]{
            count++
            i++
            if count == 9{
                break
            }
        } 
        result = result + strconv.Itoa(count) + string(str[i])
        count = 1
    }
	return result
}
