package main

import (
	"fmt"
	"ipfs-cid/internal"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cid-local <filename>")
		return
	}

	filename := os.Args[1]

	cid := internal.Cid(filename)
	fmt.Println(cid)
}
