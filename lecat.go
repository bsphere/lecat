package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/bsphere/le_go"
	"log"
	"os"
)

func main() {
	var tokenflag string

	flag.StringVar(&tokenflag, "token", "", "-token=<logentries_token>")

	flag.Parse()

	if tokenflag == "" {
		fmt.Fprintln(os.Stderr, "Usage: lecat -token=<logentries_token>")
		flag.PrintDefaults()
		os.Exit(2)
	}

	le, err := le_go.Connect(tokenflag)
	if err != nil {
		log.Fatal(err)
	}

	defer le.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		le.Println(scanner.Text())
	}
}
