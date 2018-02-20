package main

import (
	"fmt"

)

func absDifirence(a int, b int)  int{
	var res int

	if a-b < 0 {
		res = b-a
	}

	if a - b > 0 {
		res = a - b
	}
	return res

}

func main()  {
	var lenght int
	fmt.Scanf("%d", &lenght)

//1.make a 2-dimensional array
	arr := make([][]int, lenght) //make rows
	//folowed by columns
	for col := range arr {
		arr[col] = make([]int, lenght)
	}
//2.populate the array
	var num int
	for i:=0; i <lenght; i++ {
		for j := 0; j<lenght; j++ {
			fmt.Scanf("%d", &num)
			arr[i][j] = num
		}
	}

//3. Find diagonal sums

	var rightD, leftD int

	for j := 0; j < lenght; j++{
		for k := 0; k < lenght; k++ {
			if j == k {
				leftD += arr[j][k]
			}
			if j+k == lenght-1 {
				rightD += arr[j][k]
			}
		}
	}

	//fmt.Println(leftD, rightD)

	fmt.Println(absDifirence(leftD, rightD))

}
