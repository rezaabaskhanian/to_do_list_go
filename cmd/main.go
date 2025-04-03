package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rezaabaskhanian/to_do_list_go/internal/domain"
	"github.com/rezaabaskhanian/to_do_list_go/internal/usecase"
)

func main() {
	// دستورات خط فرمان
	createTaskCmd := flag.NewFlagSet("create-task", flag.ExitOnError)
	createAssigneeCmd := flag.NewFlagSet("create-assignee", flag.ExitOnError)

	// پارامترهای برای "create-task"
	taskTitle := createTaskCmd.String("title", "", "Title of the task")
	taskDescription := createTaskCmd.String("description", "", "Description of the task")

	// پارامترهای برای "create-assignee"
	assigneeName := createAssigneeCmd.String("name", "", "Name of the assignee")
	assigneeEmail := createAssigneeCmd.String("email", "", "Email of the assignee")

	// چک کردن اینکه کاربر کدام دستور را اجرا کرده است
	if len(os.Args) < 2 {
		fmt.Println("Expected 'create-task' or 'create-assignee' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-task":
		createTaskCmd.Parse(os.Args[2:])
		if *taskTitle == "" || *taskDescription == "" {
			log.Fatal("Title and description are required for creating a task.")
		}
		// ایجاد تسک
		usecase := usecase.NewTaskUsecase() // Initialize usecase with necessary dependencies
		task := domain.Task{
			Title:       *taskTitle,
			Description: *taskDescription,
		}
		createdTask, err := usecase.CreateTask(task.Title, task.Description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task created: %+v\n", createdTask)

	case "create-assignee":
		createAssigneeCmd.Parse(os.Args[2:])
		if *assigneeName == "" || *assigneeEmail == "" {
			log.Fatal("Name and email are required for creating an assignee.")
		}
		// ایجاد مسئول
		usecase := usecase.NewAssigneeUsecase() // Initialize usecase with necessary dependencies
		assignee := domain.Assignee{
			Name:  *assigneeName,
			Email: *assigneeEmail,
		}
		createdAssignee, err := usecase.CreateAssignee(assignee.Name, assignee.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Assignee created: %+v\n", createdAssignee)

	default:
		fmt.Println("Expected 'create-task' or 'create-assignee' subcommand")
		os.Exit(1)
	}
}
