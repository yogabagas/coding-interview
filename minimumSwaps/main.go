package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {

	startPoint := len(arr) / 2
	totalSwap := 0
	for i := 0; i < len(arr)-1; i++ {

		if i < startPoint {
			if arr[i] > arr[startPoint] {
				arr[i], arr[startPoint] = arr[startPoint], arr[i]
				totalSwap++
				break
			}
		} else if i > startPoint {
			if arr[i] < arr[startPoint] {
				arr[i], arr[startPoint] = arr[startPoint], arr[i]
				totalSwap++
				break
			}
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if i >= len(arr)-1 {
				return int32(totalSwap)
			}
		}
	}

	return int32(totalSwap)

}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	// nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// arrTemp := strings.Split(readLine(reader), " ")

	// var arr []int32

	// for i := 0; i < int(n); i++ {
	//     arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	//     checkError(err)
	//     arrItem := int32(arrItemTemp)
	//     arr = append(arr, arrItem)
	// }

	res := minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6})
	fmt.Println(res)

	// fmt.Fprintf(writer, "%d\n", res)

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
