package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter numbers (press Enter on empty line to calculate):")
		sum := decimal.NewFromInt(0)

		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				break
			}

			num, err := decimal.NewFromString(input)
			if err != nil {
				fmt.Println("Invalid number, try again.")
				continue
			}

			sum = sum.Add(num)
		}

		fmt.Printf("= %s\n\n", sum.StringFixed(2))
	}
}
