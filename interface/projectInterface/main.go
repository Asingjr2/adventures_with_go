package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Creating type that can be used as customer writer
type logWriter struct{}

func main() {
	url := "https://www.google.com/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error with site", err)
		os.Exit(1)
	}

	// Function creates bytes with specific size of 9999 bytes
	bs := make([]byte, 99999)
	fmt.Println(string(bs))
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}

// Simplying creating the below method converts logwriter to use writer interface
func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
