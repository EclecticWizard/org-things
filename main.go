package main

import (
	"fmt";
	"os";
	"io";
	"log";
	"encoding/csv"
)

// TODO check for "things-cli"
// If you need it. Use pipx install things-cli

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Can't find the home directory")
	}
	file, err := os.Open(home + "/Todo/things_export.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'
	lineNumber := 0
	records := make([][]string, 0)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("err while reading CSV file: %s", err)
		}
		if lineNumber == 0 {
			lineNumber++
			continue
		}
		lineNumber++
		records = append(records, record)

	}

	f, err := os.Create("things.org")
		if err != nil {
			log.Fatal(err)
		}
	defer f.Close()
	/*
	* Headers
	* "Title";"Type";"URI";"Creation Date";"Modification Date";"Due Date";"Start Date";"Completion Date";"Recurring";"Heading";"Project";"Area";"Subtask";"Notes";"Tags"

	*/

	for _, v := range records {
		if len(v) != 15 {
			continue
		}
		title, kind, uri, creationDate, modificationDate, dueDate, startDate, completionDate, recurring, heading, project, area, subtask, notes, tags := v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11], v[12], v[13], v[14]
		fmt.Printf("Title: %s Type: %s URI: %s Creation: %s  modificationDate: %s dueDate: %s startDate: %s completionDate: %s recurring: %s heading: %s project: %s area: %s subtask: %s notes: %s tags: %s  \n", title, kind, uri, creationDate, modificationDate, dueDate, startDate, completionDate, recurring, heading, project, area, subtask, notes, tags)
		if completionDate != ""{
			continue
		}
		todo := fmt.Sprintf("* TODO %s <%s> \n", title, dueDate)
		if dueDate != ""{
			todo += fmt.Sprintf("\tDEADLINE: %s\n", dueDate)
		}
		todo += fmt.Sprintf("\t:PROPERTIES:\n")
		todo += fmt.Sprintf("\tCreated: %s\n", creationDate)


		_, err := f.WriteString(todo)
		if err != nil {
			log.Fatal(err)
		}
	}
}
