package main

import (
	"fmt"
	"bufio"
	// "errors"
	"os"
	"strings"
	"example.com/note/note"
	"example.com/note/todo"

)

type saver interface {
	Save() error 
}

type displayer interface {
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }
type outputtable interface {
	saver
	displayer
}


func add[T int|float64|string](a, b T) T {
	return a + b 
}

func main() {
	result := add(1, 2.1)
	fmt.Println(result)
	printSomething(2)
	title, content:= getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}

func printSomething(data interface{}) {
	intVal, ok := data.(int)
	if ok {
		intVal += 1
		fmt.Println(intVal)
		return
	}
	floatVal, ok := data.(float64)
	if ok {
		fmt.Println(floatVal)
		return
	}
	stringVal, ok := data.(string)
	if ok {
		fmt.Println(stringVal)
		return
	}
	// switch data.(type) {
	// 	case int:
	// 		fmt.Println(data)
	// 	case float64:
	// 		fmt.Println(data)
	// 	case string:
	// 		fmt.Println(data)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	data.Save()

	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Note saved successfully")
	return nil 
}

func getNoteData() (string, string) {
	title:= getUserInput("Note title: ")
	
	content:= getUserInput("Note content: ")

	return title, content
}

// MEMBACA LONG INPUT DARI KEYBOARD , KALAU SCAN HANYA UNTUK SATU KATA
func getUserInput(promptText string) (string) {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}