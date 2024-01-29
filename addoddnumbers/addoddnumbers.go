package addoddnumbers

import "fmt"

func Addoddnumbers() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	OddNumbers:=[]int{}

	var total int=0

	for _, oddNumber := range numbers {
		if oddNumber%2 != 0 {
			total+=oddNumber		
			OddNumbers = append(OddNumbers, oddNumber)
		}
	
     
	}	
	fmt.Println(OddNumbers)
    fmt.Println("Tek sayıların toplamı",total)
}