package main

import (
	"fmt"
	"net/http"
	"io"
	"errors"
	"os"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("recieved / request\n");
	io.WriteString(w, "This is hello world")
}

func main() {
	http.HandleFunc("/", getHello);
	err := http.ListenAndServe(":8080", nil);
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server close\nd")
	} else if err != nil {
		fmt.Printf("error starting server\n", err);
		os.Exit(1)
	}
}

