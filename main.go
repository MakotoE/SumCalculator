package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter numbers (press Enter on empty line to calculate):")
		var sum float64

		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				break
			}

			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid number, try again.")
				continue
			}

			sum += num
		}

		fmt.Printf("= %.2f\n\n", sum)
	}
}
