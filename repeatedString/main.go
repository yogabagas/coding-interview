package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	// Write your code here
	// Write your code here
	var na int64
	for _, ns := range s {
		if string(ns) == "a" {
			na++
		}
	}

	lengthString := len(s)
	divideLength := n / int64(lengthString)

	multiply := divideLength * na

	modString := n % int64(lengthString)

	for i := int64(0); i < modString; i++ {
		if s[i] == 'a' {

			multiply++
		}
	}

	return multiply

}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// s := readLine(reader)

	// n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)

	result := repeatedString("abcac", 10)

	log.Println(result)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
