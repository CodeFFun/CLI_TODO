package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "reflect"
)

// type Task struct{
// 	Id int;
// 	description string;
// 	status bool;
// 	date string;
// }

func displayTask() {
	csvFile := checkFileExists()
	readFromCsvFile(csvFile)

}

func checkFileExists() *os.File {
	stat, err := os.OpenFile("file/task.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("File does not exist, creating file ...", err)
		_, err := os.Create("file/task.csv")
		if err != nil {
			panic(err)
		}
	}
	return stat
}

func writeToCsvFile(csvFile *os.File) {
	var taskList []Task = []Task{
		{Id: 3, Description: "Task w", Status: false, Date: "2021-01-11"},
	}
	// fmt.Println(taskList)

	w := csv.NewWriter(csvFile)

	for _, record := range taskList {
		recordStr := []string{
			fmt.Sprintf("%d", record.Id),
			record.Description,
			fmt.Sprintf("%t", record.Status),
			record.Date,
		}
		if err := w.Write(recordStr); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}

func readFromCsvFile(csvFile *os.File) [][]string {
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	createFormat(records)
	return records
}

func readSingleTask(csvFile *os.File, id int) []string {
	r := csv.NewReader(csvFile)
	for {
		records, err := r.Read()
		if err != nil {
			fmt.Print("Task not found")
			log.Fatal(err)
			break
		}
		if records[0] == fmt.Sprint(id) {
			return records
		}
	}
	return nil
}

func updtaeTask(csv *os.File, id int) {
	records := readSingleTask(csv, id)
	createFormat(records)
}
