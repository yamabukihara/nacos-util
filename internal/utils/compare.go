package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

func Compare(source, target string) (string, error) {

	diff := difflib.UnifiedDiff{
		A:       difflib.SplitLines(source),
		B:       difflib.SplitLines(target),
		Context: 3,
	}
	result, err := difflib.GetUnifiedDiffString(diff)
	return result, err
}

func Confirm(prompt string) bool {
	for {
		fmt.Fprintf(os.Stderr, "%s (Y/n): ", prompt)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("read user input failed: %v", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Fprintln(os.Stderr, "please input Y/y or N/n.")
			continue
		}
	}
}
