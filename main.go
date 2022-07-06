package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxSize = 5
var some = make([]string, 0)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the maximum number of notes: ")
	line, _, err := reader.ReadLine()
	if err != nil {
		return
	}

	parseInt, err := strconv.ParseInt(string(line), 10, 32)
	if err != nil {
		return
	}

	maxSize = int(parseInt)

	for {
		fmt.Print("Enter a command and data: ")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Try again")
		}

		userInput := strings.SplitN(string(line), " ", 2)

		switch userInput[0] {
		case "create":
			if len(userInput) != 2 {
				fmt.Println("[Error] Missing note argument")
			} else {
				create(userInput[1])
			}
		case "delete":
			if len(userInput) < 2 {
				fmt.Println("[Error] Missing position argument")
			} else {
				argNum, err := strconv.Atoi(userInput[1])
				if err != nil {
					fmt.Println("[Error] Invalid position: " + userInput[1])
				} else if argNum < 1 || argNum > maxSize {
					fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", argNum, maxSize)
				} else {
					deleteCmd(argNum)
				}
			}
		case "update":
			if len(userInput) == 1 {
				fmt.Println("[Error] Missing position argument")
			} else {

				var updateUserInput = strings.SplitN(userInput[1], " ", 2)
				switch len(updateUserInput) {
				case 1:
					fmt.Println("[Error] Missing note argument")
				case 2:
					argNum, err := strconv.Atoi(updateUserInput[0])
					if err != nil {
						fmt.Println("[Error] Invalid position: " + updateUserInput[0])
					} else if argNum < 1 || argNum > maxSize {
						fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", argNum, maxSize)
					} else {
						update(argNum, updateUserInput[1])
					}
				}
			}

		case "clear":
			clear()
		case "list":
			list()
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func deleteCmd(argNum int) {
	if len(some) < argNum {
		fmt.Println("[Error] There is nothing to delete")
	} else {
		some = append(some[:argNum-1], some[argNum:]...)
		fmt.Println("[OK] The note at position " + strconv.Itoa(argNum) + " was successfully deleted")
	}
}

func update(argNum int, data string) {
	if len(some) < argNum {
		fmt.Println("[Error] There is nothing to update")
	} else {
		some[argNum-1] = data
		fmt.Println("[OK] The note at position " + strconv.Itoa(argNum) + " was successfully updated")
	}
}

func clear() {
	some = make([]string, 0)
	fmt.Println("[OK] All notes were successfully deleted")
}

func list() {
	if len(some) == 0 {
		fmt.Println("[Info] Notepad is empty")
	} else {
		for i, s := range some {
			fmt.Printf("[Info] %d: %s\n", i+1, s)
		}
	}
}

func create(data string) {
	if len(some) < maxSize {
		some = append(some, data)
		fmt.Println("[OK] The note was successfully created")
	} else {
		fmt.Println("[Error] Notepad is full")
	}
}
