package friendnumber




import "fmt"


func FriendNumberFind(){
	number1:=220
	number2:=280
	total1:=0
	total2:=0
	

	for i:=1;i<number1;i++{
		if number1%i==0{
			total1=total1+i
		}
	}
	for i:=1;i<number2;i++{
		if number2%i==0{
			total2=total2+i
		}
	}
	if total1==number2 && total2==number1{
		fmt.Println(number1,"and",number2,"are friend numbers")
	}else{
		fmt.Println(number1,"and",number2,"are not friend numbers")
	}

}