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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	fmt.Println(convertToBin(n))
	getMaximunOnes(convertToBin(n))
}

func convertToBin(i int32) string {
	i64 := int64(i)
	return strconv.FormatInt(i64, 2) // base 2 for binary
}

func getMaximunOnes(value string) {

	var iMaxQuantity = 0
	var iQuantity = 0

	for i := 0; i < len(value); i++ {
		if value[i] == '1' {
			iQuantity++
		} else {
			if iQuantity > iMaxQuantity {
				iMaxQuantity = iQuantity
			}
			iQuantity = 0
		}
	}

	if iQuantity > iMaxQuantity {
		iMaxQuantity = iQuantity
	}

	fmt.Println(iMaxQuantity)
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
