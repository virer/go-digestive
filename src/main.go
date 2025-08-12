package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// The hashDigest function performs the core logic of the dictionary attack for a single password.
func hashDigest(username, realm, password, method, uri, nonce, nc, cnonce, qop, expectedResponse string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Calculate HA1
	a1 := fmt.Sprintf("%s:%s:%s", username, realm, strings.TrimSpace(password))
	hA1 := md5.Sum([]byte(a1))
	HA1 := hex.EncodeToString(hA1[:])

	// Calculate HA2
	a2 := fmt.Sprintf("%s:%s", method, uri)
	hA2 := md5.Sum([]byte(a2))
	HA2 := hex.EncodeToString(hA2[:])

	// Calculate the final response hash
	response2 := fmt.Sprintf("%s:%s:%s:%s:%s:%s", HA1, nonce, nc, cnonce, qop, HA2)
	hResponse := md5.Sum([]byte(response2))
	finalResponse := hex.EncodeToString(hResponse[:])

	// Compare with the expected response
	if finalResponse == expectedResponse {
		resultChan <- password
	}
}

func main() {
	// Define command-line flags
	username := flag.String("username", "", "Username")
	wordlistPath := flag.String("wordlist", "", "Path to the wordlist")
	method := flag.String("method", "", "HTTP method")
	nonce := flag.String("nonce", "", "nonce")
	cnonce := flag.String("cnonce", "", "cnonce")
	uri := flag.String("uri", "", "uri")
	qop := flag.String("qop", "", "qop")
	response := flag.String("response", "", "response")
	nc := flag.String("nc", "", "nc")
	realm := flag.String("realm", "", "realm")
	flag.Parse()

	// Validate required flags
	if *username == "" || *wordlistPath == "" || *method == "" || *nonce == "" || *cnonce == "" || *uri == "" || *qop == "" || *response == "" || *nc == "" || *realm == "" {
		fmt.Println("All arguments are required.")
		flag.Usage()
		os.Exit(1)
	}

	// Open the wordlist file
	file, err := os.Open(*wordlistPath)
	if err != nil {
		log.Fatalf("Failed to open wordlist file: %v", err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	resultChan := make(chan string, 1)

	// Create a buffered scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := scanner.Text()
		wg.Add(1)
		// Launch a goroutine for each password
		go hashDigest(*username, *realm, password, *method, *uri, *nonce, *nc, *cnonce, *qop, *response, resultChan, &wg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading wordlist: %v", err)
	}

	// We are done launching goroutines, so we can wait for them to finish.
	// This goroutine waits for all workers to finish and then closes the channel.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// The program will now wait for a value to be received from the channel.
	// This is where we print the password immediately.
	foundPassword, ok := <-resultChan
	if ok {
		fmt.Println("Username =", *username)
		fmt.Println("Password =", strings.TrimSpace(foundPassword))
		os.Exit(0)
	} else {
		// If the channel is closed and no password was found, we will get here.
		fmt.Println("Password not found.")
		os.Exit(1)
	}
}
