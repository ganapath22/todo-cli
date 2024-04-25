package todo

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func (todo *Todo) Validate() error {
	if len(todo.Title) == 0 {
		return errors.New("title is empty")
	}

	if len(todo.Desc) > 100 {
		return errors.New("description is empty")
	}

	return nil

}

func (todo *Todo) Print() {
	fmt.Printf("------------\n")
	doneYesNo := "No"
	if todo.Done {
		doneYesNo = "Yes"
	}
	fmt.Printf("Title: %v \nDescription: %v\nCreated At: %v\nDone: %v\n",
		todo.Title, todo.Desc, todo.CreatedAt.Format("2006-01-02 15:04"), doneYesNo)
	fmt.Printf("------------\n")
}
