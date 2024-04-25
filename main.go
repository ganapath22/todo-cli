package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ganapath22/todo-cli/todo"
)

func main() {
	fmt.Printf("This is a CLI app to manage simple todo lists. It is written to learn and explore Go. \n")

	t := todo.Todo{
		Title:     "First list",
		Desc:      "Functionality check",
		Done:      true,
		CreatedAt: time.Now(),
	}

	if t.Validate() != nil {
		fmt.Printf("\nTodo is invalid. Error: %v\n", t.Validate())
	} else {
		t.Print()
	}

	todoBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("\nsomething went wrong while encoding to JSON: %v\n", err)
	}

	todoJSONString := string(todoBytes)
	fmt.Printf("\nJSON String: %v\n", todoJSONString)
	fmt.Printf("\nJSON Bytes: %v\n", []byte(todoJSONString))

	

}
