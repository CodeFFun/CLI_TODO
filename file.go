package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
    "regexp"
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
    _, err := os.Stat("./file") 
    if os.IsNotExist(err){
    if err != nil{
        err := os.Mkdir("./file", 0750)
        if err != nil {
            fmt.Println("Error creating directory:", err)
            return nil
        }
    }
}
	stat, _ := os.OpenFile("file/task.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
        fmt.Println("Error opening the file", err)
		fmt.Println("File does not exist, creating file ...", err)


		stat, err := os.Create("file/task.csv")
		if err != nil {
        fmt.Println(err)
		panic(err)
		}
        return stat
	}
	return stat
}

func writeToCsvFile(csvFile *os.File) {
	var taskList []Task = []Task{
		{Id: 1, Description: "Task 1", Status: false, Date: "2021-01-11"},
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

func updateDesc(file *os.File, index int){
}

func updateDate(file *os.File, index int){

}

func update(file *os.File, index int, slug string, data ...string) {
    //func update(){
    record := readSingleTask(file, index)
    r,_ := regexp.MatchString(`[0-9]{4}(-|\/)(0[0-9]|1[0-2])(-|\/)([0-2][0-9]|3[0-2])`, "2024-08-12")
    fmt.Println(r)


}
