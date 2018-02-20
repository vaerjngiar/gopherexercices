package main

import "fmt"

func main()  {
	var point, aliceRes, bobRes int

	alice := make([]int,3)
	bob := make([]int, 3)

	for i:=0; i <3; i++ {
		fmt.Scanf("%d", &point)
		alice[i] = point
	}

	for i:=0; i <3; i++ {
		fmt.Scanf("%d", &point)
		bob[i] = point
	}

	//fmt.Println(alice)
	//fmt.Println(bob)

	for i:=0; i <3; i++ {
		switch  {
		case alice[i] > bob[i]:
			aliceRes++
		case bob[i] > alice[i]:
			bobRes++
		}
	}

	fmt.Println(aliceRes, bobRes)
}
