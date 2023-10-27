package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define a struct to hold the student information
type Student struct {
	Name    string
	Age     string
	Subject string
}

// Function to parse a line based on the delimiter
func parseLine(line string) Student {
	parts := strings.Split(line, ",")
	if len(parts) < 3 {
		panic("Invalid line: " + line)
	}
	return Student{
		Name:    strings.TrimSpace(parts[0]),
		Age:     strings.TrimSpace(parts[1]),
		Subject: strings.TrimSpace(parts[2]),
	}
}

func main() {
	// Open the flat file
	file, err := os.Open("students.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		student := parseLine(line)
		fmt.Printf("Name: %s, Age: %s, Subject: %s\n", student.Name, student.Age, student.Subject)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
