package main

import "myworkshop/restfullapi"

/*
    "myworkshop/channels"
	"fmt"
	"myworkshop/bignumberfind"
	"myworkshop/findnumbergame"
	"myworkshop/friendNumber"
	"myworkshop/primenumbers"
	"myworkshop/addoddnumbers"
    "myworkshop/dictionary"
	"myworkshop/varargs"
	"myworkshop/goroutines"
	"time"
*/

func main(){
	/*
fmt.Println("*************Prime Number Find*************")
primenumbers.PrimeNumberFind(7)
fmt.Println("*************Big Number Find*************")
bignumberfind.BigNumberWrite(50,20,5)
fmt.Println("*************Find Number Game*************")
findnumbergame.FindNumber()
fmt.Println()
fmt.Println("*************Friend Number Find*************")
friendnumber.FriendNumberFind()
addoddnumbers.Addoddnumbers()
dictionary.Dictionary()
fmt.Println("*************Varargs Sum*************")
result:=varargs.VarargsSum(1,2,3,4,5,6,7,8,9,10)
fmt.Println("Sayıların Toplamı",result)
// Goroutine Example 1
go goroutines.EvenNumbers()
go goroutines.OddNumbers()
time.Sleep(5*time.Second)
fmt.Println("Main bitti")
//Channel Example
evenNumberCn:=make(chan int)
oddNumberCn:=make(chan int)
go channels.EvenNumbers(evenNumberCn)
go channels.OddNumbers(oddNumberCn)

evenNumberTotal,oddNumberTotal:=<- evenNumberCn,<- oddNumberCn
multiply:=evenNumberTotal*oddNumberTotal
fmt.Println("Çarpım:", multiply)




*/
restfullapi.JsonPost()



}