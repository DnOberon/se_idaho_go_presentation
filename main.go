package main

// package names become the accessor prefix for utilizing their contents
import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.txt") // os.Create returns a pointer to a os.File type
	if err != nil {
		log.Fatal(err) // Go's best practices specify that errors should be handled, not panic'd
	}

	for {
		addressesAndInterfaces(file)

		time.Sleep(5 * time.Second) // time.Second is a constant provided by the time package
	}
}

func addressesAndInterfaces(file *os.File) {
	// research io.Reader and io.Writer interfaces - it is better to use those interfaces than using WriteString() all over
	file.WriteString(time.Now().String() + "\n")

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
			log.Printf("%s : %s", inf.Name, a.String())

			file.WriteString(fmt.Sprintf("%s : %s\n", inf.Name, a.String()))
		}

	}

}
