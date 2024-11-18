package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"
)

type Task struct {
	Id          int
	Description string
	Status      bool
	Date        string
}

func putIntoTask(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Task Description: ")
	description, _ := reader.ReadString('\n')
	fmt.Print("Enter a Task Date: ")
	date, _ := reader.ReadString('\n')
	task := Task{Id: 1, Description: description, Status: false, Date: date}
	fmt.Sprintln("Task created:", task)
}

func createFormat(arr interface{}){
	switch records := arr.(type){
	case [][]string:
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 15, ' ', 0)
		fmt.Fprintln(w, "ID\tDescription\tCompleted\tDate\t")
		for _, record := range records {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", record[0], record[1], record[2], record[3])
		}
		w.Flush()	
	case []string:
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 15, ' ', 0)
		fmt.Fprintln(w, "ID\tDescription\tCompleted\tDate\t")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", records[0], records[1], records[2], records[3])
		w.Flush()
	default:
		fmt.Println("Invalid type")
	}
}
