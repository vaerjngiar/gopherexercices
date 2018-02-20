package main

import (
	"fmt"
	"strings"
)

func main()  {
	var input string

	fmt.Scanf("%s\n", &input)
	result := 1

	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			result++
		}
	}

	fmt.Println(result)

}
