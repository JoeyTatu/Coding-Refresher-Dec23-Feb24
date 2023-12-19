package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// fmt.Printf("domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecord\nPlease enter email or domain to check, or type 'quit' to quit:\n")

	fmt.Printf("Please enter email or domain to check, or type 'quit' to quit:\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		input := getUserInput(scanner)

		if input == "quit" || input == "exit" {
			os.Exit(2)
			break
		}

		checkDomain(input)
	}
}

func getUserInput(scanner *bufio.Scanner) string {
	fmt.Print("> ")
	scanner.Scan()
	return strings.ToLower(strings.TrimSpace(scanner.Text()))
}

func checkDomain(input string) {
	email := input
	var domain string

	parts := strings.Split(email, "@")
	if len(parts) > 1 {
		domain = parts[1]
	} else {
		domain = input
	}

	hasMX, hasSPF, hasDMARC := false, false, false
	spfRecord, dmarcRecord := "Not found", "Not found"

	lookupMXRecords(domain, &hasMX)
	lookupTXTRecords(domain, "v=spf1", &hasSPF, &spfRecord, "TXT")
	lookupTXTRecords("_dmarc."+domain, "v=DMARC1", &hasDMARC, &dmarcRecord, "DMARC")

	printDomainInfo(domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	main()
}

func lookupMXRecords(domain string, hasMX *bool) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records: %v\n", err)
		return
	}

	if len(mxRecords) > 0 {
		fmt.Print("MX Record: ")
		for _, mx := range mxRecords {
			fmt.Printf("%s ", mx.Host)
		}
		fmt.Println()
		*hasMX = true
	}
}

func lookupTXTRecords(domain, prefix string, flag *bool, record *string, recType string) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records: %v\n", err)
		return
	}
	for _, txtRec := range txtRecords {
		if strings.HasPrefix(txtRec, prefix) {
			fmt.Printf("%v Record: %v\n", recType, txtRec)
			*flag = true
			*record = txtRec
		}
	}
}

func printDomainInfo(domain string, hasMX, hasSPF bool, spfRecord string, hasDMARC bool, dmarcRecord string) {
	fmt.Printf("\n\nDomain: %v\nHasMX: %v\nHasSPF: %v\nSPF Record: %v\nHasDMARC: %v\nDMARC Record: %v\n",
		domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	if hasMX && hasSPF && hasDMARC {
		fmt.Printf("**********\n%v domain is set up for emails.\n**********\n\n", domain)
	} else {
		fmt.Printf("**********\n%v domain is NOT set up for emails.\n**********\n\n", domain)
	}
}
