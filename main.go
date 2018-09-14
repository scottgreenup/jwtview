package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Pulled from dgrijalva/jwt-go
func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}

func main() {

	header := flag.Bool("header", false, "If specified will show the header")
	claims := flag.Bool("claims", false, "If specified will show the claims")
	signature := flag.Bool("signature", false, "If specified will show the signature")

	flag.Parse()

	// Read in the JWT from stdin
	reader := bufio.NewReader(os.Stdin)
	line := make([]byte, 0, 0)

	for {
		linePrefix, prefix, readErr := reader.ReadLine()
		line = append(line, linePrefix...)

		if prefix != true {
			break
		}

		if readErr != nil {
			fmt.Printf("Error reading line: %v\n", readErr)
		}
	}

	// Process the JWT
	encodedSegment := strings.Split(string(line), ".")

	if len(encodedSegment) != 3 {
		fmt.Printf("Expected 3 parts, found %d parts\n", len(encodedSegment))
		return
	}

	// We only want to decode the first 2 segments, leave the signature (3rd segment) as base64
	var decodedSegment [2]string
	for i := 0; i < 2; i++ {
		decoded, decodeErr := DecodeSegment(encodedSegment[i])

		if decodeErr != nil {
			fmt.Printf("Error decoding an encoded part: %v\n", decodeErr)
			return
		}

		decodedSegment[i] = string(decoded)
	}

	showAll := !(*header || *claims || *signature)

	// Output the JWT segments
	if showAll {
		fmt.Println(decodedSegment[0])
		fmt.Println(decodedSegment[1])
		fmt.Println(encodedSegment[2])
		return
	}

	if *header {
		fmt.Println(decodedSegment[0])
	}

	if *claims {
		fmt.Println(decodedSegment[1])
	}

	if *signature {
		fmt.Println(encodedSegment[2])
	}
}
