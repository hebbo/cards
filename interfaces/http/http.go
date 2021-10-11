package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func (l LogWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("custom implementation")

	return len(b), nil
}

func Call() {
	fmt.Println("Reaching out to Google...")
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	b := make([]byte, 100)
	_, err = resp.Body.Read(b)

	if err != nil {
		fmt.Println("Error reading body")
		os.Exit(1)
	}

	writter := LogWriter{}

	io.Copy(writter, resp.Body)
}
