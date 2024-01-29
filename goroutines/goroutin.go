package goroutines

import (
	"fmt"
	"time"
)

func EvenNumbers(){
	for i:=0;i<10;i++{
		if i%2==0{
			fmt.Println("Çift sayı: ",i)
			time.Sleep(1*time.Second)
		}

	}
}
func OddNumbers(){
	for i:=0;i<10;i++{
		if i%2!=0{
			fmt.Println("Tek sayı: ",i)
			time.Sleep(1*time.Second)
		}

	}
}
