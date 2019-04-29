package main

import (
	"log"
	"os"
	"time"
)

// Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.
type Outputter interface {
	Output([]string)
}

// simple struct for log output
type LogOutputter struct {
}

func (lo LogOutputter) Output(output []string) {
	for _, line := range output {
		log.Print(line)
	}
}

type FileOutputter struct {
	// as long as the struct satisfies the interface, we can add as many additional methods and properties as needed
	file *os.File
}

func (fo FileOutputter) Output(output []string) {
	fo.file.WriteString(time.Now().String() + "\n")

	for _, line := range output {
		fo.file.WriteString(line + "\n")
	}
}
