package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	trie "../trie"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr,
			"Provide the app with the compiled dict.")
		os.Exit(1)
	}

	dictPath := flag.Arg(0)
	dict, err := trie.LoadTrie(dictPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		distance, word := fields[1], fields[2]
		answers := dict.SearchCloseWords(word, distance)
		fmt.Println(prettyPrint(answers))
	}
}
