package channels

func EvenNumbers(evenNumberCn chan int) {
	total := 0
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			total = total + i
		}

	}
	evenNumberCn <- total

}
func OddNumbers(oddNumbersCn chan int) {

	total := 0
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			total = total + i

		}

	}
	oddNumbersCn <- total
}