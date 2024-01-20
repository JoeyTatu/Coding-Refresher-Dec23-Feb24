package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	HTML string
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		HTML: `!DOCTYPE html>
        <html lang="en" data-theme="auto">
        <head>
            <title>403 - Forbidden</title>
			<link rel="icon" type="image/png" href="https://www.pngall.com/wp-content/uploads/2016/06/6a00d83451b36c69e20168eaa60893970c-600wi.png">
    </head>
        </head>
        <body style="background-color: black; color: white;">
            <center>
                <h1>
                    <img src="https://www.pngall.com/wp-content/uploads/2016/06/6a00d83451b36c69e20168eaa60893970c-600wi.png" alt="Blocked icon image">
                    <br>This site is blocked for your safety, or due to ad-blocking
                </h1>
            </center>
        </body>
        </html>
		`,
	}

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(page.HTML))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on http://127.0.0.1:80")

	log.Fatal(http.ListenAndServe(":80", nil))
}
