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
	// 检查必要的环境变量
	if os.Getenv("LARK_API_KEY") == "" {
		fmt.Fprintln(os.Stderr, "错误: 未设置 LARK_API_KEY 环境变量")
		fmt.Fprintln(os.Stderr, "请参考 README.md 或 .env.example 文件进行配置")
		fmt.Fprintln(os.Stderr, "提示: 运行 'source .env' 加载环境变量")
		os.Exit(1)
	}

	fmt.Println("🎉 Go 智能体已启动！")
	fmt.Println("📖 输入您的任务，或输入 'exit' 或 'quit' 退出")
	fmt.Println("💡 支持的工具：bash、read_file、write_file、edit_file")
	fmt.Println()

	hist := []anthropic.MessageParam{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" || input == "exit" || input == "quit" {
			break
		}

		_, output, err := Chat(input, hist)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ 错误: %v\n", err)
			fmt.Print("> ")
			continue
		}

		fmt.Println(output)
		fmt.Print("> ")
	}

	fmt.Println()
	fmt.Println("👋 再见！")
}
