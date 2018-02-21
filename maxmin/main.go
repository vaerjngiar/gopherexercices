package main

import "fmt"

func main()  {
	var num, total  int64
	arr := make([]int64, 5)

	for i:= 0; i <5; i++ {
		fmt.Scanf("%d", &num)
		arr[i] = num
	}

	//fmt.Println(arr)
	var max = arr[0]
	var min = arr[0]

	for j := 0; j < len(arr); j++ {
		total += arr[j]
		if arr[j] < min {min = arr[j]}
		if arr[j] > max {max = arr[j]}
	}

	fmt.Println(total-max, total-min)
}
