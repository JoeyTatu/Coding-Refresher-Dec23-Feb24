package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	var status string

	connection, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable.\nError: %v", destination, err)
	} else {
		defer connection.Close()

		status = fmt.Sprintf("[UP] %v is reachable.\n From: %v\n To: %v", destination, connection.LocalAddr(), connection.RemoteAddr())
	}

	return status
}

func main() {
	fmt.Print("Enter the domain (e.g., 'google.com'):\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	domain := strings.TrimSpace(scanner.Text())

	port := "8080"

	fmt.Println("Checking with port 8080:")
	status := check(domain, port)
	if strings.Contains(status, "DOWN") {
		fmt.Printf("FAILED: %v is not reachable with port 8080.\n", domain)

		port = "80"

		fmt.Println("\nChecking with port 80:")
		status = check(domain, port)
	}

	fmt.Println(status)
}
