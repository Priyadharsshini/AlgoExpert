package main

func CommonCharacters(strings []string) []string {
    finalmap := make(map[string]int)
    tempmap := make(map[string]int)
    var result []string
	for _, char := range strings[0]{
        if _, ok := finalmap[string(char)] ; !ok {
            finalmap[string(char)] = 1
        }
    }
    for i := 1; i < len(strings); i++ {
        for _, char := range strings[i]{
            if _, ok := tempmap[string(char)] ; !ok {
                tempmap[string(char)] = 1
            }
        }
        for key, _ := range tempmap{
            if _, ok := finalmap[key]; ok{
                finalmap[key]++
            }
        }
        for key := range tempmap {
            delete(tempmap, key)
        }
    }
    for key, value := range finalmap{
        if value == len(strings){
            result = append(result, key)
        }
    }
    if len(result) == 0{
        return []string{}
    }
	return result
}
