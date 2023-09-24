package main

import (
	// "bufio"
	// "fmt"
	// "io"
	"bufio"
	"io"
	"log"
	"os"
)

// An helper function that panics when the results are nil.
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    CHANGELOG := "/home/myData/Code/SumahoNikki/CHANGELOG.md"
    // Put all the file into memory; not the best way to do this
    data, err := os.ReadFile(CHANGELOG)
    check(err)
    log.Print(string(data))

    log.Print("READING BYTES\n")

    // A better way is to obtain a file and then process it
    file, err := os.Open(CHANGELOG)
    check(err)

    // Reading part of the file
    bytes5 := make([]byte, 5)
    result1, err := file.Read(bytes5)
    check(err)
    log.Printf("%d bytes: %s\n", result1, string(bytes5[:result1]))

    log.Print("Seeking and reading \n")

    // Reading a specific part
    offset1, err := file.Seek(6,0)
    check(err)

    bytes2 := make([]byte, 2)
    result2, err := file.Read(bytes2)
    check(err)
    log.Printf("%d bytes @ %d: ", result2, offset1)
    log.Printf("%v\n", string(bytes2[:result2]))


    log.Print("READING USING IO")
    offset2, err := file.Seek(6,0)
    check(err)
    result3, err := io.ReadAtLeast(file, bytes5, 2)
    check(err)
    log.Printf("%d bytes @ %d: %s\n", result3, offset2, string(bytes5))

    log.Printf("REWINDING THE FILE")
    _, err = file.Seek(0,0)
    check(err)

    result4 := bufio.NewReader(file)
    bytes4, err := result4.Peek(50)
    check(err)
    log.Printf("4 bytes: %s\n", string(bytes4))


}
