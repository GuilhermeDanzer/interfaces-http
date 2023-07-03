package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type longWriter struct{}

func main() {

	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := longWriter{}
	io.Copy(lw, resp.Body)
}

func (longWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
