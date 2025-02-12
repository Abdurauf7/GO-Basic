package main

import (
	"bufio"
	"example/section6/note"
	"example/section6/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Please enter todo:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Node title:")
	content := getUserInput("Node content:")
	return title, content
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Print("Saving node failed")
		return err
	}

	fmt.Println("Saving node success")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

// if we want to display any value we can add simply interface{} or any type
func printSomething(value interface{}) {
	// switch value.(type) { // defining type of value string or interger or any type
	// case int:
	// 	fmt.Println("This is an integer")
	// case string:
	// 	fmt.Println("This is a string")
	// case float64:
	// 	fmt.Println("This is a float")
	// default:
	// 	fmt.Println("Unknown type")

	// }
	// fmt.Println(value)

	// Another way to define type
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("This is an integer", typedVal)
		return
	}

}
