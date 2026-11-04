package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// io.Reader
// 1. The io.Reader has one method - Read
// 2. Signature -

type Reader interface {
	Read(p []byte) (int, error)
}

// 3. Takes a byte slice to read the data into - where the data read from the source is stored
// 4. Returns two values -
// a. First is the number of bytes read from the source
// b. Second is the error

// 5. The caller of Read should always process n bytes returned before considering error
// 6. This way we can correctly handle I/O errors that happen after reading some bytes
// 7. Sentinel error io.EOF is returned when EOF is reached

// Common Implementations
// 1. *os.File - used to read data from a file
// 2. strings.Reader - creates a new reader from the provided string
// 3. http.Request.Body - allows you to supply the request body through reader

func readerExample() {
	f, err := os.Open("letters.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	n, err := countAlphabets(f) // we can pass f as os.File implements io.Reader
	if err != nil {
		panic(err)
	}

	fmt.Println("Letters read from a file:", n)

	s := strings.NewReader("Hello Worldddd")

	num, err := countAlphabets(s)
	fmt.Println("Letters read from a string:", num)
}

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)

		for _, l := range buffer[:n] {
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}

		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return 0, err
		}
	}
}

// io.Writer
// 1. It is used to write data to a data stream
// 2. Takes the source data as slice of byte
// 3. Signature -

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Implementing io.Writer
// 1. If written bytes are less than the provided bytes you should return an error
// 2. Provided slice of byte is a source of writing and should not be modified

// Common Implementations
// 1. *os.File - for writing to files
// 2. os.Stdout - for writing to standard out
// 3. http.ResponseWriter - for writing response body of http request

func writerExample() {
	f, err := os.Create("writing.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	n, err := writeString("Hello World!!!", f) // we can pass f as it satisfies the interface
	if err != nil {
		panic(err)
	}

	fmt.Println("Written Bytes ", n)
}

func writeString(s string, w io.Writer) (int, error) {
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {
	readerExample()
	writerExample()
}

// Summary
// 1. The io.Reader allows you to read data from source file, network connections etc
// 2. The io.Writer allows you to write data to file, network connections etc
// 3. Allow an abstraction view of destination - make code generic & re-usable
