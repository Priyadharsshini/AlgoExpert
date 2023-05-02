package main

func BalancedBrackets(s string) bool {
    var a []string
	if len(s) == 0 {
        return false
    }
    for _, c := range s {
        val := string(c)
        if len(a) == 0 && (val == "}" || val == ")" || val == "]" ){
            return false
        }
        if val == "{" || val == "(" || val == "["{
              a = append(a,val)  
        } else if val == "}" {
            if a[len(a)-1] == "{" {
                 a = a[0:len(a)-1]
            } else {
                return false
            }
        } else if val == ")"  {
            if a[len(a)-1] == "(" {
                 a = a[0:len(a)-1]
            } else {
                return false
            }
        } else if val == "]" {
            if a[len(a)-1] == "[" {
              a = a[0:len(a)-1]
            } else {
                return false
            }
        } 
    }
    if len(a) == 0 {
        return true
    }
	return false
}
