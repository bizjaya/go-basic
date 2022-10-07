package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID		  string   `json:"id"`
	Item 	  string   `json:"item"`
	Completed bool     `json:"completed"`
}

var todos = []todo{
	{ ID: "1", Item: "Clean Room", Completed: false },
	{ ID: "2", Item: "Read Book", Completed: false },
	{ ID: "3", Item: "Record Video", Completed: false },

}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}

func addTodo(context *gin.Context){
	var newTodo todo
	if err := context.BindJSON(&newTodo);  err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err !=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
	}

	context.IndentedJSON(http.StatusOK, todo)


}

func toggleTodoStatus(context *gin.Context){

	id := context.Param("id")
	todo, err := getTodoById(id)

	if err !=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
	}


	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)

}





func getTodoById(id string) (*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")

}

func EncodeJson() {
	
	TODOS := []todo {
		{ ID: "1", Item: "Clean Room", Completed: false },
		{ ID: "2", Item: "Read Book", Completed: false },
		{ ID: "3", Item: "Record Video", Completed: false },
	}

	finalJson, err := json.MarshalIndent(TODOS, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
			"id": "66",
			"item": "do things",
			"website": Completed,
		}
	`)
    
	var TODOZ todo

	isValidJson := json.Valid(jsonDataFromWeb)

	if isValidJson {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &TODOZ)
	}else{
		fmt.Println("JSON IS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)


	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T", k, v ,v)
	}
}

func main() {
    router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)


	router.Run("localhost:9090")
	
	var test = "gsheet"
	fmt.Print("Hlello ", test, "w orld")
}


