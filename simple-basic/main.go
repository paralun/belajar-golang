package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-basic/model"
	"simple-basic/service"
	"strconv"
)

func main() {
	svc := service.NewTaskService()
	scanner := bufio.NewScanner(os.Stdin)
	coba := true
	for coba {
		fmt.Println("Task Application :\n" +
			"1. New Task\n" +
			"2. View Task\n" +
			"3. Exit")
		fmt.Printf("Choose [1-3]: ")
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())

		switch input {
		case 1:
			fmt.Print("Task:  ")
			scanner.Scan()
			data := scanner.Text()
			task := model.Task{ToDoItem:data}
			svc.New(task)
		case 2:
			tasks := svc.View()
			fmt.Println("-----------")
			fmt.Println("  TASKS")
			fmt.Println("-----------")
			for _, v := range tasks {
				fmt.Println(v.ToDoItem)
			}
			fmt.Println("-----------")
		case 3:
			coba = false
		default:
			fmt.Println("Action belum dipilih")
		}
	}

	fmt.Println("Selesai")
}