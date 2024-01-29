package primenumbers


import "fmt"



func PrimeNumberFind(number int){	

	isPrime:=true

	for i:=2;i<number;i++{

		if number %i==0{
			isPrime=false
		}
	}
	if isPrime==true{
		fmt.Println(number,"number is prime number")
	}else{
		fmt.Println(number,"number is not a prime number")
	}
	
}


