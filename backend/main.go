package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Backend...")

	http.ListenAndServe(":8080", nil)
}
