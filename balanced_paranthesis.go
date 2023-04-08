package main

import "fmt"

var a []string

func paranthesisChecker(s string) {
	if len(s)%2 != 0 {
		fmt.Println("Invalid")
		return
	}
	valid := true
	for _, val := range s {
		if string(val) == "(" || string(val) == "{" || string(val) == "[" {
			a = append(a, string(val))
		} else if string(val) == ")" && a[len(a)-1] == "(" {
			a = a[0 : len(a)-1]
		} else if string(val) == "}" && a[len(a)-1] == "{" {
			a = a[0 : len(a)-1]
		} else if string(val) == "]" && a[len(a)-1] == "[" {
			a = a[0 : len(a)-1]
		} else {
			valid = false
			break
		}
	}
	if valid == true {
		fmt.Println("Yes valid paranthesis")
	} else {
		fmt.Println("Invalid")
	}

}

func main() {
  s := "{}[]()"
	paranthesisChecker(s)

}
