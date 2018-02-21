package main

import "fmt"

func main()  {

	var n int
	fmt.Scanf("%d", &n)

	for i:=0; i <n; i++{
		fmt.Printf("%*s", n-i, "#")
		for j:= 0; j <i; j++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
