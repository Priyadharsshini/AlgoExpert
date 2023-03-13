package main

func MinimumCharactersForWords(words []string) []string {
    hm1 := make(map[rune]int)
    final := make(map[rune]int)
    output := make([]string, 0, 10)
	for _, word := range words{
        for _, letter := range word{
            if _, ok := hm1[letter]; ok{
                hm1[letter]++
            } else {
                 hm1[letter] = 1
            }
        }
        for k, v := range hm1 {
            if final[k] < v{
            final[k] = v
            }
        }
      hm1 = make(map[rune]int)


    }
    for key, value := range(final){
        for i := 1 ; i <= value; i++{
             output = append(output,string(key))   
        }
    }
    return output
}
