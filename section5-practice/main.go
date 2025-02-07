package main

import (
	"bufio"
	"example/practice/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Print("Saving node failed")
		return
	}

	fmt.Println("Saving node success")
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
