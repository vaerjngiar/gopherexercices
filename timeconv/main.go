package main

import (
	"fmt"
	"strconv"
	//"time"
)

func main()  {
	var timeString string
	fmt.Scanf("%s", &timeString)
	hours, _ := strconv.Atoi(timeString[:2])
	minutes, _ := strconv.Atoi(timeString[3:5])
	seconds, _ := strconv.Atoi(timeString[6:8])
	meridiem := timeString[8:10]


	//fmt.Println(hours, minutes, seconds, meridiams)
	if meridiem == "PM" && hours < 12 {
		hours += 12
	}

	if meridiem == "PM" && hours == 12 {
		hours = 0
	}



	fmt.Printf("%02d:%02d:%02d\n", hours, minutes, seconds)

}
