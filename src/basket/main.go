package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	min := scores[0]
	max := scores[0]
	cmin, cmax := int32(0), int32(0)
	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			cmax++
			max = scores[i]
		}
		if scores[i] < min {
			cmin++
			min = scores[i]
		}
	}
	return []int32{cmax, cmin}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer func() {
		if terr := stdout.Close(); terr != nil {
			fmt.Print(terr.Error())
		}
	}()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	scoresTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var scores []int32

	for i := 0; i < int(n); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	result := breakingRecords(scores)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	if err := writer.Flush(); err != nil {
		fmt.Print(err.Error())
	}
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
