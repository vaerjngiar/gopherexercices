package main

import (
	"fmt"
	"math/big"
)

func main()  {
	var num1, num2 int64
	fmt.Scanf("%d",&num1)
	fmt.Scanf("%d",&num2)

	t1 := big.NewInt(num1)
	t2 := big.NewInt(num2)

	var n int
	fmt.Scanf("%d",&n)

	t3 := big.NewInt(0)

	for i:= 2; i < n; i++{


		t3 = add(t1, mul(t2,t2))
		t1 = t2
		t2 = t3
		//fmt.Println(t1,t2,t3)
	}

	fmt.Println(t2)




}

func add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

