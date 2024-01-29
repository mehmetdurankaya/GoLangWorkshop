package bignumberfind

import (
"fmt"
)


func BigNumberWrite(a, b, c int) {

	if a > b && a > c {
		fmt.Println("En büyük sayı: ",a)
	} else if b> c {
		fmt.Println("En büyük sayı: ",b)
	} else {
		fmt.Println("En büyük sayı: ",c)
	}
}