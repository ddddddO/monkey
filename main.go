package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ddddddO/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hellll, %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
