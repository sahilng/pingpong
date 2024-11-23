package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ANSI escape codes for colors
const (
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func main() {
	// Flag for hiding the HTML output
	hideHTML := flag.Bool("h", false, "Hide the HTML output")
	flag.Parse()

	// Check if URL is provided
	args := flag.Args() // Get non-flag arguments
	if len(args) < 1 {
		fmt.Println("Usage: pingpong [-h] <url>")
		return
	}
	urlStr := args[0]

	// Add "http://" if no protocol is specified
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "http://" + urlStr
	}

	timeout := time.Duration(10 * time.Second) // 10 seconds timeout
	client := http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow redirects but capture when they happen
			return nil
		},
	}

	// Start measuring the request time
	start := time.Now()
	resp, err := client.Get(urlStr)
	duration := time.Since(start) // Stop the timer immediately after the request completes

	if err != nil {
		fmt.Printf("Error fetching the URL: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the HTTP status is in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("HTTP request failed with status: %s\n", resp.Status)
		return
	}

	// Detect if there was a redirect
	finalURL := resp.Request.URL.String()
	wasRedirected := finalURL != urlStr

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response: %s\n", err)
		return
	}

	// Print the HTML output unless -h is specified
	if !*hideHTML {
		fmt.Println(string(body))
	}

	// Print request duration and redirect warning (if applicable)
	fmt.Printf("%s%.0f%s%sms%s\n", Green, duration.Seconds()*1000, Reset, Blue, Reset) // Number in green, "ms" in blue
	if wasRedirected {
		fmt.Printf("%sRedirect detected: Request was redirected to %s. For more accurate timing, try testing directly with https.%s\n", Yellow, finalURL, Reset)
	}
}
