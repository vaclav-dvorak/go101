package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	if n < 0 || n > 100000 {
		panic("incorrect input")
	}

	book := make(map[string]string)

	for i := 0; i < int(n); i++ {
		line := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		book[line[0]] = line[1]
	}

	c := 0
	for {
		if c > 100000 {
			panic("too many queries")
		}
		query := readLine(reader)
		c++
		if len(query) != 0 {
			if book[query] != "" {
				fmt.Printf("%s=%s\n", query, book[query])
			} else {
				fmt.Println("Not found")
			}
		} else {
			break
		}
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
