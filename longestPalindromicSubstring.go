package main

func LongestPalindromicSubstring(str string) string {
	longestPalindrome := ""
    for i, _ := range str {
        left := i-1
        right := i+1
        for (left >=0 && right < len(str) && str[left] == str[right]){
            left --
            right ++
        }
        if right-left-1 > len(longestPalindrome){
                longestPalindrome = str[left+1:right]
        }
        // For even length 
        left = i-1
        right = i
        for (left >=0 && right < len(str) && str[left] == str[right]){
               left --
                right ++ 
        }
        if right-left-1 > len(longestPalindrome){
                longestPalindrome = str[left+1:right]
        }
    }    
        
    
	return longestPalindrome
}
    
    
