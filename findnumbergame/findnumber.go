package findnumbergame

import "fmt"

func FindNumber() {

	keptNumber:= 80
	predictedNumber := 0

	fmt.Println("guess a number")
	fmt.Scanln(&predictedNumber)
	count := 1
	for keptNumber != predictedNumber {

		if predictedNumber < keptNumber {
			fmt.Println("a larger number")
			fmt.Scanln(&predictedNumber)
		}
		if predictedNumber > keptNumber {
			fmt.Println("a smaller number")
			fmt.Scanln(&predictedNumber)
		}
		count++
	}

	fmt.Printf("Bravo you got in on your %dth guess! ", count)

}