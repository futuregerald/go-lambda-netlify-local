package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Take advantage of features built in to third party routers like method filtering, param parsing, etc.
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/.netlify/functions/testfunc2", lambdaHandler)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":3000", r))
	} else {
		log.Fatal(gateway.ListenAndServe(":3000", r))
	}
}

func lambdaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Looks like you made it to testfunc2!"))
}
