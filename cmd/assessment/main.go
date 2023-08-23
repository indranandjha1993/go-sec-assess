package main

import (
	"flag"
	"fmt"
	"github.com/indranandjha1993/go-sec-assess/pkg/headerchecker"
	"github.com/indranandjha1993/go-sec-assess/pkg/tlschecker"
	"github.com/indranandjha1993/go-sec-assess/pkg/urlchecker"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net/url"
	"strings"
)

const (
	colorReset       = "\033[0m"
	colorRed         = "\033[31m"
	colorGreen       = "\033[32m"
	colorYellow      = "\033[33m"
	colorBlue        = "\033[34m"
	colorPurple      = "\033[35m"
	colorCyan        = "\033[36m"
	colorWhite       = "\033[37m"
	checkMark        = "\u2714" // Unicode for check mark symbol
	crossMark        = "\u2718" // Unicode for cross mark symbol
	missingIndicator = "[MISSING]"
	presentIndicator = "[PRESENT]"
)

func main() {
	host := flag.String("host", "", "The URL to assess (e.g., https://example.com)")
	flag.Parse()

	if err := validateURL(*host); err != nil {
		log.Fatalf("%v\nUsage: ./main -host <URL_TO_ASSESS>", err)
	}

	displayHeader(*host)

	parsedURL, err := url.Parse(*host)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	domain := parsedURL.Hostname()

	tlsVersion, cipher := tlschecker.CheckTLS(domain)
	displayTLSInfo(tlsVersion, cipher)

	headerResults, err := headerchecker.CheckHeaders(*host)
	if err != nil {
		log.Fatalf("Error checking headers: %v", err)
	}

	for _, result := range headerResults {
		if result.Value == "" {
			displayHeaderCheck(result.Name, missingIndicator)
		} else {
			displayHeaderCheck(result.Name, result.Value)
		}
	}
}

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", padding)
}

func displayHeader(url string) {
	terminalWidth, _, _ := terminal.GetSize(0)

	fmt.Println(strings.Repeat("=", terminalWidth))
	fmt.Println(centerText("[ Security Assessment ]", terminalWidth))
	fmt.Println(centerText(fmt.Sprintf("Target: %s", url), terminalWidth))
	fmt.Println(strings.Repeat("=", terminalWidth))
}

func displayTLSInfo(version, cipher string) {
	prefix := "[+] Checking TLS Information..."
	fmt.Println(prefix)

	// Handle TLS Version
	if version == "" {
		fmt.Printf("    |- TLS Version: %s%s%s %s%s%s\n", colorRed, "Error", colorReset, colorRed, crossMark, colorReset)
	} else {
		fmt.Printf("    |- TLS Version: %s%s%s\n", colorCyan, version, colorReset)
	}

	// Handle Cipher Suite
	if cipher == "Unknown or uncommon cipher suite" || cipher == "" {
		fmt.Printf("    |- Cipher Suite: %s%s%s %s%s%s\n", colorRed, "Error", colorReset, colorRed, crossMark, colorReset)
	} else {
		fmt.Printf("    |- Cipher Suite: %s%s%s %s%s%s\n", colorGreen, cipher, colorReset, colorGreen, checkMark, colorReset)
	}
}

func displayHeaderCheck(header, value string) {
	if value == missingIndicator {
		fmt.Printf("    |- %s %s%s%s\n", header, colorRed, crossMark, colorReset)
	} else {
		fmt.Printf("    |- %s: %s%s%s %s%s%s\n", header, colorCyan, value, colorReset, colorGreen, checkMark, colorReset)
	}
}

func validateURL(url string) error {
	if url == "" {
		return fmt.Errorf("please provide a URL using the -url flag")
	}

	if !urlchecker.IsValidURL(url) {
		return fmt.Errorf("provided URL (%s) is not valid or well-formed", url)
	}

	return nil
}
