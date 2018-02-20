package main

import "fmt"

func main()  {
	var lenght int
	var num int64
	var res int64
	fmt.Scanf("%d", &lenght)
	//fmt.Println(lenght)
	arr := make([]int64, lenght)

	for i := 0; i <lenght; i++ {
		fmt.Scanf("%v", &num)
		//fmt.Println(num)
		arr[i] = num
		res += num
	}

	//fmt.Println(arr)
	fmt.Println(res)


}
