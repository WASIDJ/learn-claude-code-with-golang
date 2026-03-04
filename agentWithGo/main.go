package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var client = anthropic.NewClient(
	option.WithAPIKey(os.Getenv("LARK_API_KEY")),
	option.WithBaseURL("https://ark.cn-beijing.volces.com/api/coding"),
)

func main() {
	hist := []anthropic.MessageParam{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" || input == "exit" || input == "quit" {
			break
		}

		hist, output, err := Chat(input, hist)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		_ = hist

		fmt.Println(output)
		fmt.Print("> ")
	}
}
