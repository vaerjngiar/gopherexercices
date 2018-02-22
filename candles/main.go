package main

import "fmt"

func main()  {
	var n int
	fmt.Scanf("%d", &n)

	var tall, freq, height  int

	for i :=0; i < n; i++ {
		//var height int
		fmt.Scanf("%d", &height)
		if height > tall {
			tall = height
			freq = 1
		} else if height == tall {freq ++}

	}

	fmt.Println(freq)
}
