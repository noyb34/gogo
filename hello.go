package main

import "fmt"
import "os"

func main() {

	if len(os.Args) > 1 {

		fmt.Printf("Hello, world.\n")

	} else {

		fmt.Printf("Hello, P. Roch!\nThis is a test.\n")

	}

}
