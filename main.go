package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	myFigure := figure.NewFigure("Email checker", "", true)
	fmt.Println(strings.Repeat("#", 100))
	fmt.Println("\n")
	myFigure.Print()
	fmt.Println("\n")
	fmt.Println(strings.Repeat("#", 100))
	fmt.Print("\n Enter your email: ")

	for scanner.Scan() {
		email := scanner.Text()
		domain := extractDomain(email)
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input", err)
	}
}

func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "Invalid email format"
	}
	return parts[1]
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check for MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
	} else {
		hasMX = len(mxRecords) > 0
	}

	// Check for SPF records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// Check for DMARC records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	red := "\033[31m"  // ANSI escape code for red color
	reset := "\033[0m" // ANSI escape code to reset color
	// Output formatting
	fmt.Println("\n")
	fmt.Println(strings.Repeat("#", 35))
	fmt.Printf("# %sProvided Domain: %s%s\n", red, reset, domain)
	fmt.Println(strings.Repeat("#", 35))
	fmt.Printf("# %sOutput:%s\n", red, reset)
	fmt.Printf("# MX Records: %s\n", getPresenceString(hasMX))
	fmt.Printf("# SPF Records: %s\n", getPresenceString(hasSPF))
	fmt.Printf("# DMARC Records: %s\n", getPresenceString(hasDMARC))
	fmt.Printf("# SPF Record: %s\n", spfRecord)
	fmt.Printf("# DMARC Record: %s\n", dmarcRecord)
	fmt.Println(strings.Repeat("#", 35))
}

func getPresenceString(present bool) string {
	if present {
		return "Present"
	}
	return "Absent"
}
