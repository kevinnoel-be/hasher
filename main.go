package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kevinnoel-be/hasher/pkg/hash"
)

type args struct {
	iterations  int
	salt        string
	privateSalt string
}

func main() {
	a := args{}
	flag.IntVar(&a.iterations, "iterations", 1, "Number of iterations")
	flag.StringVar(&a.salt, "salt", "", "Base 64 encoded salt (public)")
	flag.StringVar(&a.privateSalt, "private-salt", "", "Base 64 encoded private salt")
	flag.Parse()

	var password []byte

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			password = append(password, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Failed to read pipe data: %v", err)
		}
	} else {
		fmt.Print("Data to hash: ")

		var pass string
		_, _ = fmt.Scanf("%s", &pass)
		password = []byte(pass)
	}

	if len(password) == 0 {
		log.Fatal("Empty data")
	}

	computedHash, _ := hash.Compute(hash.Request{
		Data:        password,
		Salt:        b64decode(a.salt),
		PrivateSalt: b64decode(a.privateSalt),
		Iterations:  a.iterations,
	})

	fmt.Printf("%v\n", base64.StdEncoding.EncodeToString(computedHash))
}

func b64decode(s string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("Failed to base64 decode: %v", err)
	}
	return decoded
}
