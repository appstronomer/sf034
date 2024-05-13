package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	logStderr := log.New(os.Stderr, "", 0)

	// соблюдаю пожелание 2
	if len(os.Args) < 3 {
		logStderr.Println("please suply two arguments: input file name; output file name")
		os.Exit(1)
	}
	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	// соблюдаю пожелание 1 в части 'regexp'
	re := regexp.MustCompile(`(?m)^(\s*(\d+)\s*([+-]{1})\s*(\d+)\s*=\s*)\?\s*$`)
	// соблюдаю пожелание 1 в части 'ioutil', хотя пакет устарел и правильнее использовать 'os'.
	// А в целом лучше читать построчно через 'bufio.Reader.ReadString()', решать каждый пример
	// сразу после чтения, и записывать через 'bufio.Writer.WriteString()' сразу после решения -
	// так можно обрабатывать большие файлы
	inputBytes, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		logStderr.Printf("error on input file read: %v\n", err.Error())
		os.Exit(1)
	}

	// соблюдаю пожелание 3, передав "os.O_CREATE"
	// соблюдаю пожелание 4, передав "os.O_TRUNC"
	outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logStderr.Printf("error on output file write: %v\n", err.Error())
		os.Exit(1)
	}
	defer outputFile.Close()

	// соблюдаю пожелание 5
	outputFileWriter := bufio.NewWriter(outputFile)
	defer outputFileWriter.Flush()

	for _, item := range re.FindAllStringSubmatch(string(inputBytes), -1) {
		expression := item[1]
		operation := item[3]
		firstOperand, _ := strconv.Atoi(item[2])
		secondOperand, _ := strconv.Atoi(item[4])
		var result int

		switch operation {
		case "+":
			result = firstOperand + secondOperand
		case "-":
			result = firstOperand - secondOperand
		default:
			logStderr.Printf("operation not supported: %v\n", operation)
		}

		_, err = outputFileWriter.WriteString(fmt.Sprintf("%v%v\n", expression, result))
		if err != nil {
			logStderr.Printf("error on output file write: %v\n", err.Error())
			os.Exit(1)
		}
	}
}
