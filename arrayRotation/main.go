package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func rotLeft(a []int32, d int32) []int32 {
	// Write your code here
	var newArrays = []int32{}

	for j := 0; j < int(d); j++ {
		firstNum := a[0]
		for i := 0; i < len(a)-1; i++ {
			a[i] = a[i+1]

			if i+1 == len(a)-1 {
				a[len(a)-1] = firstNum
			}
			newArrays = a
		}
	}

	return newArrays
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// d := int32(dTemp)

	// aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// var a []int32

	// for i := 0; i < int(n); i++ {
	// 	aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
	// 	checkError(err)
	// 	aItem := int32(aItemTemp)
	// 	a = append(a, aItem)
	// }

	a := []int32{1, 2, 3, 4, 5, 6, 7}

	result := rotLeft(a, 2)

	log.Println(result)

	// for i, resultItem := range result {
	// 	fmt.Fprintf(writer, "%d", resultItem)

	// 	if i != len(result)-1 {
	// 		fmt.Fprintf(writer, " ")
	// 	}
	// }

	// fmt.Fprintf(writer, "\n")

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
