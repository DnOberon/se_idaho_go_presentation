package main

// package names become the accessor prefix for utilizing their contents
import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	var logout LogOutputter
	var fileout FileOutputter

	file, err := os.Create("log.txt") // os.Create returns a pointer to a os.File type
	if err != nil {
		log.Fatal(err) // Go's best practices specify that errors should be handled, not panic'd
	}

	fileout.file = file

	go func(outputters ...Outputter) { // anon func can be declared inline - or you can do something like go yourfunc()
		for {
			addressesAndInterfaces(outputters...)

			time.Sleep(5 * time.Second) // time.Second is a constant provided by the time package
		}
	}(&logout, &fileout) // in order to maintain scope, specify and pass in needed variables

	// go supports closures - here we use it to specify which outputter the handler should use to call
	// the last seen function
	http.HandleFunc("/", outputterHandle(&fileout))

	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func outputterHandle(outputter Outputter) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")

		json.NewEncoder(resp).Encode(outputter.LastSeen())
	}
}

func addressesAndInterfaces(outputters ...Outputter) { // variadic function parameters are treated as an array
	output := map[string][]string{}

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	for _, inf := range interfaces {
		addresses, err := inf.Addrs()
		if err != nil {
			log.Printf("%s", err.Error())
			continue
		}

		for _, a := range addresses {
			output[inf.Name] = append(output[inf.Name], a.String()) // slices can be appended to
		}

	}

	for _, outputter := range outputters {
		outputter.Output(output)
	}
}
