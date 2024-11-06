package main

import (
	"bufio"
	. "fmt"
	"os"
	. "strconv"
	. "strings"
)

func main() {
	// 	write your code here
	var size int
	Print("Enter the maximum number of notes: ")
	Scan(&size)

	memoryNotes := make([]string, 0, size)

	scanner := bufio.NewScanner(os.Stdin)
	Print("\nEnter a command and data: ")

	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := Split(line, " ")
		command := lineSlice[0]
		data := Join(lineSlice[1:], " ")

		switch command {
		case "create":
			if len(memoryNotes) >= cap(memoryNotes) {
				Println("[Error] Notepad is full")
			} else if line == "create" {
				Println("[Error] Missing note argument")
			} else {
				memoryNotes = append(memoryNotes, addToNotes(data))
				Println("[OK] The note was successfully created")
			}
		case "list":
			if len(memoryNotes) == 0 {
				Println("[Info] Notepad is empty")
			} else {
				Println(listNotes(memoryNotes))
			}
		case "update":
			if line == "update" {
				Println("[Error] Missing position argument")
			} else {
				updateReq := Split(data, " ")
				updatePos, err := Atoi(updateReq[0])
				updateData := Join(updateReq[1:], " ")
				if err != nil {
					Printf("[Error] Invalid position: %v\n", updateReq[0])
				} else if updateData == "" {
					Println("[Error] Missing note argument")
				} else if updatePos > cap(memoryNotes) || updatePos < 1 {
					Printf("[Error] Position %d is out of the boundaries [1, %d]\n", updatePos, cap(memoryNotes))
				} else if len(memoryNotes) == 0 || updatePos > len(memoryNotes) {
					Println("[Error] There is nothing to update\n")
				} else {
					updateNote(memoryNotes, updatePos-1, updateData)
					Printf("[OK] The note at position %d was successfully updated\n", updatePos)
				}

			}
		case "delete":
			if line == "delete" {
				Println("[Error] Missing position argument")
			} else {
				pos, err := Atoi(data)
				if err != nil {
					Printf("[Error] Invalid position: %v\n", data)
				} else if pos < 1 || pos > cap(memoryNotes) {
					Printf("Position %d is out of the boundaries [1, %d]\n", pos, cap(memoryNotes))
				} else if len(memoryNotes) == 0 || pos > len(memoryNotes) {
					Println("[Error] There is nothing to delete")
				} else {
					memoryNotes = deleteNote(memoryNotes, pos-1)
					Printf("[OK] The note at position %d was successfully deleted\n", pos)
				}
			}
		case "clear":
			memoryNotes = memoryNotes[0:0]
			Println("[OK] All notes were successfully deleted")
		case "exit":
			Println("[Info] Bye!")
			return
		default:
			Println("[Error] Unknown command")
		}

		Print("\nEnter a command and data: ")
	}
}

func addToNotes(data ...string) string {
	var b Builder
	for _, bit := range data {
		b.WriteString(bit)
	}
	return b.String()
}

func listNotes(data []string) string {
	var result []string
	for i, piece := range data {
		log := Sprintf("[Info] %d: %s", i+1, piece)
		result = append(result, log)
	}
	return Join(result, "\n")
}

func updateNote(notes []string, index int, data ...string) {
	for i := range notes {
		if i == index {
			notes[i] = addToNotes(data...)
		}
	}
}

func deleteNote(notes []string, index int) []string {
	if index == 0 {
		return append(notes[0:0], notes[1:]...)
	} else if index == len(notes)-1 {
		return notes[:len(notes)-1]
	} else {
		return append(notes[:index], notes[index+1:]...)
	}
}
