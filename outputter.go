package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.
type Outputter interface {
	Output(map[string][]string)
	LastSeen() map[string][]string
}

// simple struct for log output
type LogOutputter struct {
	Last map[string][]string
}

func (lo *LogOutputter) Output(output map[string][]string) {
	for inf, val := range output {

		for _, line := range val {
			log.Printf("%s : %s", inf, line)
		}

	}

	lo.Last = output
}

func (lo LogOutputter) LastSeen() map[string][]string {
	return lo.Last
}

type FileOutputter struct {
	Last map[string][]string
	file *os.File
}

func (fo *FileOutputter) Output(output map[string][]string) {
	fo.file.WriteString(time.Now().String() + "\n")

	for inf, val := range output {

		for _, line := range val {
			fo.file.WriteString(fmt.Sprint("%s : %s", inf, line))
		}

	}

	fo.Last = output
}

func (fo FileOutputter) LastSeen() map[string][]string {
	return fo.Last
}
