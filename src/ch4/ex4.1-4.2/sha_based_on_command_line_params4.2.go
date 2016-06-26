package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

func main() {
	sha := flag.String("sha", "256", "Set the length of hash. Can be set to 256, 384 or 512")
	flag.Parse()

	content := strings.Join(flag.Args(), "")
	fmt.Println(content)

	bytes := []byte(content)

	switch *sha {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256(bytes))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384(bytes))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(bytes))
	default:
		fmt.Printf("%s is not a valid option: It should be 256 | 384 | 512.\n", *sha)
	}
}
