package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JoeyTatu/coffee-api/db"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

// Global port variable for both Main and Serve
var port = os.Getenv("PORT") // 8080

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	return srv.ListenAndServe()
}

func main() {
	// fmt.Println("USER:", os.Getenv("USER"))
	// fmt.Println("PW:", os.Getenv("PW"))

	var cfg Config
	cfg.Port = port

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
