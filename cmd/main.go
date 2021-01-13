package main

import (
	"api-customers/application/router"
	"api-customers/application/tracer"
	"fmt"
	"log"
	"net/http"
)

func main() {

	flush := tracer.InitTracer()
	defer flush()

	port := 8080

	http.Handle("/", router.Router())

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening in port: %d\n", port)
}
