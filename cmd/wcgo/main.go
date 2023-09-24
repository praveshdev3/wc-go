package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var printBytesCount, printCharCount, printLineCount, printWordsCount bool
	var filename string
	flag.BoolVar(&printBytesCount, "b", false, "print bytes count")
	flag.BoolVar(&printCharCount, "c", false, "print char count")
	flag.BoolVar(&printLineCount, "l", false, "print line count")
	flag.BoolVar(&printWordsCount, "w", false, "print word count")
	flag.Parse()
	filename = flag.CommandLine.Arg(0)
	if !printBytesCount && !printCharCount && !printLineCount && !printWordsCount {
		printLineCount = true
		printWordsCount = true
		printBytesCount = true
	}
	var file []byte
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(1)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		file, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(2)
		}
	} else {
		file, err = os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(3)
		}
	}
	if printBytesCount {
		fmt.Printf("Bytes => %d\n", len(file))
	}
	if printCharCount {
		chars := len(bytes.Runes(file))
		fmt.Printf("Chars => %d\n", chars)
	}
	if printLineCount {
		lines := 0
		for _, b := range file {
			if b == '\n' {
				lines++
			}
		}
		fmt.Printf("Lines => %d\n", lines)
	}
	if printWordsCount {
		words := len(bytes.Fields(file))
		fmt.Printf("Words => %d\n", words)
	}
}
