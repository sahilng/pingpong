/*
   pingpong.go: A simple utility to measure the request time of a given URL and display the response.
   I used ChatGPT by OpenAI to write this code... wow!
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ANSI escape codes for colors
const (
	Green = "\033[32m"
	Blue  = "\033[34m"
	Reset = "\033[0m"
)

func main() {
	// Flag for hiding the request time
	hideTime := flag.Bool("h", false, "Hide the request time")
	flag.Parse()

	// Check if URL is provided
	args := flag.Args() // Get non-flag arguments
	if len(args) < 1 {
		fmt.Println("Usage: pingpong [-h] <url>")
		return
	}
	url := args[0]

	timeout := time.Duration(10 * time.Second) // 10 seconds timeout
	client := http.Client{
		Timeout: timeout,
	}

	// Start measuring the request time
	start := time.Now()
	resp, err := client.Get(url)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response: %s\n", err)
		return
	}

	// Calculate request duration
	duration := time.Since(start)

	fmt.Println(string(body))
	if !*hideTime {
		fmt.Printf("%s%.0f%s%sms%s\n", Green, duration.Seconds()*1000, Reset, Blue, Reset) // Number in green, "ms" in blue
	}
}
