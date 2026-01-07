package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func readInt32(scanner *bufio.Scanner) int32 {
	for {
		scanner.Scan()
		intStr := scanner.Text()
		value, err := strconv.ParseInt(intStr, 10, 32)
		if err != nil {
			fmt.Println("Not a 32 bit number please enter a valid input.")
			continue
		}
		return int32(value)
	}
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
