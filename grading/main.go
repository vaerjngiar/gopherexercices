package main

import "fmt"

func main()  {
	var n, newGrade int
	fmt.Scanf("%d", &n)
	for i:=0; i < n; i++ {
		var grade int
		fmt.Scanf("%d", &grade)



		newGrade = ((grade / 5) +1)*5
		if newGrade < 40 {
			fmt.Println(grade)
			continue
		}

		if newGrade - grade < 3 {
			fmt.Println(newGrade)
		} else {
			fmt.Println(grade)
		}


	}
}