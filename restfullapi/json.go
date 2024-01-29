package restfullapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Copletted bool   `json:"completed"`
}

func JsonGet() {
	response,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")//GET 
		if err!=nil{
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes,_:=ioutil.ReadAll(response.Body)
	bodyString:=string(bodyBytes)
	fmt.Println("Bytten Stringe çevrilen veri",bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes,&todoResponse)
	fmt.Println(todoResponse)
}
func JsonPost() {

	todo:=Todo{1,2,"Alışveriş yapılacak",false}
	jsonTodo,err:=json.Marshal(todo)
	if err!=nil{
		fmt.Println(err)
	}
	response,err:=http.Post("https://jsonplaceholder.typicode.com/todos",
	"application/json;charset=utf-8",bytes.NewBuffer(jsonTodo))//POST
	if err!=nil{
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes,_:=ioutil.ReadAll(response.Body)

	bodyString:=string(bodyBytes)
	fmt.Println(bodyString)


	var todoRes Todo
	json.Unmarshal(bodyBytes,&todoRes)
	fmt.Println(todoRes)
}