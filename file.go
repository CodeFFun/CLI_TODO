package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
    "regexp"
	// "reflect"
)


func displayTask() {
	csvFile := checkFileExists()
	readFromCsvFile(csvFile)
}

//check for dir and file and create if it does not exist. Returns the file location
func checkFileExists() *os.File {
	//check for directory
    _, err := os.Stat("./file") 
    if os.IsNotExist(err){
    if err != nil{
		//create directory
        err := os.Mkdir("./file", 0750)
        if err != nil {
            fmt.Println("Error creating directory:", err)
            return nil
        }
    }
}
	//check if file exists
	stat, _ := os.OpenFile("file/task.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
        fmt.Println("Error opening the file", err)
		fmt.Println("File does not exist, creating file ...", err)
		//creating file
		stat, err := os.Create("file/task.csv")
		if err != nil {
        fmt.Println(err)
		panic(err)
		}
        return stat
	}
	return stat
}

//takes file location and writes to the file
func writeToCsvFile(csvFile *os.File) {
	var taskList []Task = []Task{
		{Id: 1, Description: "Task 1", Status: false, Date: "2021-01-11"},
	}
	// fmt.Println(taskList)

	w := csv.NewWriter(csvFile)
	//loop until tasklist ends
	for _, record := range taskList {
		//format the record to string
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

//takes file location and reads from the whole file
func readFromCsvFile(csvFile *os.File) [][]string {
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	createFormat(records)
	return records
}

//takes file location and reads a single task

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

func updateDesc(record []string, desc string){
	record[1] = desc
}

func updateDate(record []string, date string){
	record[3] = date
}

func update(file *os.File, index int, slug string, data ...string) {
    //func update(){
    record := readSingleTask(file, index)
	if len(data) > 1{
		updateDesc(record, data[0])
		updateDate(record, data[1])
	} 
     if slug == "desc"{
		updateDesc(record, data[0])
	 } else {
		r,_ := regexp.MatchString(`[0-9]{4}(-|\/)(0[0-9]|1[0-2])(-|\/)([0-2][0-9]|3[0-2])`, "2024-08-12")
		if r == true {
			updateDate(record, data[1])
		} 
			log.Fatal("Invalid date format")
	 }

}
