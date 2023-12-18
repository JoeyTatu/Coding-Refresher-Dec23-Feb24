package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joeytatu/go-postgres/router"
	_ "github.com/lib/pq"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}
