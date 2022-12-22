package util

import (
	"strconv"
)

func ToAscii(input string) (asci string, asciInt int, err error) {
	for i := 0; i < len(input); i++ {
		asciInt += int(input[i])
		asci += strconv.Itoa(int(input[i]))
	}

	return asci, asciInt, err
}
