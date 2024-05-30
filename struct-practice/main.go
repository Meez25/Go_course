package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("HELLO")
	title, content := getNoteData()
	text := getUserInput("Todo text ")

	todo, err := todo.New(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)

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
		fmt.Println(err)
		return
	}
	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Saving the stuff succeeded !")
	return nil
}

func printSomething(value any) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Int: ", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Int: ", floatVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float64:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (string, string) {
	note := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return note, content
}
