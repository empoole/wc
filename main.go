package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	var byteFlag = flag.Bool("c", false, "returns the byte count of the given file")
	var lineFlag = flag.Bool("l", false, "returns the line count of the given file")
	var wordFlag = flag.Bool("w", false, "returns the word count of the given file")
	var charFlag = flag.Bool("m", false, "returns the character count of the given file")
	flag.Parse()
	
	var noFlags = !(*byteFlag || *lineFlag || *wordFlag || *charFlag)

	filePath := flag.Args()[0]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	bytes, lines, words, chars := count(file)

	if(*byteFlag || noFlags) {
		fmt.Print(bytes)
		fmt.Print(" ")
	}

	if(*lineFlag || noFlags) {
		fmt.Print(lines)
		fmt.Print(" ")
	}

	if(*wordFlag || noFlags) {
		fmt.Print(words)
		fmt.Print(" ")
	}

	if(*charFlag) {
		fmt.Print(chars)
		fmt.Print(" ")
	}

	fmt.Print(filePath)

}

func count(file *os.File) (byteCount int, lineCount int, wordCount int, charCount int) {
	byteCount = 0
	lineCount = 0
	wordCount = 0
	charCount = 0

	buffer := make([]byte, 1024)
	newLineChar := []byte{'\n'}
	spaceChar := []byte{' '}

	for {
		count, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if count > 0 {
			byteCount += count
			lineCount += bytes.Count(buffer[:count], newLineChar)
			wordCount += bytes.Count(buffer[:count], spaceChar)
			charCount += len(buffer[:count])
		}
	}

	return	
}
