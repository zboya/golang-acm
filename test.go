package main

import "fmt"

func main() {

	var n int //n组数据
	var str string
	var result string

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		result = result + sort(&str) + "\n"
	}
	fmt.Println(result)
}

func sort(str *string) string {
	strByte := []byte(*str)
	n := len(strByte)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if strByte[j+1] < strByte[j] {
				strByte[j+1], strByte[j] = strByte[j], strByte[j+1]
			}
		}

	}
	result := insertSpace(&strByte)
	return string(result)
}

func insertSpace(src *[]byte) []byte {
	l := len(*src)
	dst := make([]byte, 2*l)
	for i := 0; i < len(dst); i++ {
		if i%2 == 0 {
			dst[i] = (*src)[i/2]
		} else {
			dst[i] = ' '
		}
	}
	return dst
}
