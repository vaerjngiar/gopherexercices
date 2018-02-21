package main

import "fmt"

func main()  {
	var lenght, num int
	var pcount, mcount, zcount int
	fmt.Scanf("%d", &lenght)

	//arr := make([]int, lenght)

	for i :=0; i <lenght; i++ {
		fmt.Scanf("%d", &num)
		if num > 0 {pcount++}
		if num < 0 {mcount++}
		if num == 0 {zcount++}

	}

	fmt.Println(float64(pcount) / float64(lenght))
	fmt.Println(float64(mcount) / float64(lenght))
	fmt.Println(float64(zcount) / float64(lenght))


}
