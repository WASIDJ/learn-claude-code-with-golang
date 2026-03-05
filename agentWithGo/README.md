# Go 智能体项目

这是一个基于 Go 语言和火山引擎 Coding API 构建的智能体项目。智能体可以与用户进行自然语言交互，并通过执行代码和操作文件来完成任务。

## 功能特性

- 🤖 **自然语言交互**：支持与用户进行多轮对话
- 🔧 **工具链**：
  - `bash`：执行 shell 命令
  - `read_file`：读取文件内容
  - `write_file`：写入文件内容
  - `edit_file`：编辑文件内容
- 🛡️ **安全保障**：限制工作目录，防止路径逃逸；阻止危险命令
- ⚡ **实时反馈**：工具调用过程实时显示，输出截断以避免溢出

## 快速开始

### 1. 环境准备

1. 确保你已经安装了 Go 1.16+
2. 获取火山引擎 Coding API 密钥

### 2. 配置环境变量

```bash
# 复制并编辑环境变量配置
cp .env.example .env

# 编辑 .env 文件，替换 API 密钥
# 将 your_api_key_here 替换为你的火山引擎 Coding API 密钥
```

### 3. 加载环境变量

```bash
source .env
```

### 4. 启动智能体

```bash
# 构建项目
go build -o bin/agent

# 启动智能体
./bin/agent
```

### 5. 使用智能体

启动后，你会看到提示符 `> `，可以开始与智能体对话。

```
> 你好，我想创建一个简单的 Go 程序
...
```

## 支持的命令

### bash 工具

执行 shell 命令，支持读取和写入操作。

```
Example:
我需要查看当前目录下的文件
> ls -la
```

### read_file 工具

读取文件内容。

```
Example:
读取 main.go 文件的前 50 行
> cat main.go
```

### write_file 工具

写入文件内容。

```
Example:
创建一个 hello.go 文件
> write_file hello.go package main; import "fmt"; func main() { fmt.Println("Hello, World!") }
```

### edit_file 工具

编辑文件内容。

```
Example:
修改 main.go 文件
> edit_file main.go "package main" "package mypackage"
```

## 安全说明

项目实现了以下安全措施：

- 限制命令执行范围，阻止危险命令（如 `rm -rf /`, `sudo`, `shutdown` 等）
- 限制工作目录，防止路径逃逸
- 命令执行超时控制（120秒）
- 输出截断（50KB）

## 项目结构

```
agentWithGo/
├── agent.go          # 智能体核心逻辑
├── config.go         # 配置和工具定义
├── main.go           # 主入口
├── tools.go          # 工具实现
├── bin/              # 编译后的可执行文件
├── .env.example      # 环境变量配置模板
├── hello_developer.md # 开发人员欢迎说明
└── README.md         # 项目说明文档
```

## 开发说明

### 依赖管理

使用 Go Modules 管理依赖：

```bash
# 下载依赖
go mod tidy

# 升级依赖
go get -u ./...
```

### 调试模式

可以通过在 `main.go` 中添加日志输出来调试：

```go
import "log"

func main() {
	log.SetFlags(log.Lshortfile)
	// ...
}
```

## 常见问题

1. **Q: 如何更改工作目录？**
   A: 项目默认使用当前工作目录，你可以修改 `tools.go` 中的 `WORKDIR` 变量。

2. **Q: 如何添加新工具？**
   A: 在 `config.go` 中添加工具定义，在 `tools.go` 中实现执行逻辑。

3. **Q: 为什么命令执行失败？**
   A: 可能是命令超时（超过 120 秒）或被安全策略阻止。

## 许可证

[MIT License](https://opensource.org/licenses/MIT)