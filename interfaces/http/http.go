package http

import (
	"fmt"
	"net/http"
	"os"
)

func Call() {
	fmt.Println("Reaching out to Google...")
	resp, err := http.Get("https://duckduckgo.com/")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	b := make([]byte, 10000)
	_, err = resp.Body.Read(b)

	if err != nil {
		fmt.Println("Error reading body")
		os.Exit(1)
	}

	fmt.Print(string(b))
}
