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
	const LABELWEIRD string = "Weird"
	const LABELNOTWEIRD string = "Not Weird"

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	if (N % 2) != 0 {
		fmt.Println(LABELWEIRD)
	} else {
		if N >= 2 && N <= 5 {
			fmt.Println(LABELNOTWEIRD)
		} else {
			if N >= 6 && N <= 20 {
				fmt.Println(LABELWEIRD)
			} else {
				fmt.Println(LABELNOTWEIRD)
			}

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
