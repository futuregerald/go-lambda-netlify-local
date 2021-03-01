package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
)

func main() {
	http.HandleFunc("/", helloFunc)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":3000", nil))
	} else {
		log.Fatal(gateway.ListenAndServe(":3000", nil))
	}
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is go1"))
}
