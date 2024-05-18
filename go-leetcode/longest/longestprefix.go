package main

import (
	"fmt"
	"strings"
)

//program to compute the longest string prefix
/*
*

Time complexity : O(S) , where S is the sum of all characters in all strings.

In worst case if all the strings are same then 0(n)

space : o(1)

 */


func main() {

	s := []string{"flower", "flow", "flight"}
	prefix := s[0]
	for i := 1; i < len(s); i++ {
		for !strings.HasPrefix(s[i], prefix) {
			prefix = prefix[:len(prefix)-1] //flowe
		}
	}
	fmt.Println(prefix)

}
