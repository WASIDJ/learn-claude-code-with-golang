---
name: mcp-builder
description: 构建 MCP (Model Context Protocol) 服务器
---

# MCP 服务器构建技能

## 功能说明
- 快速生成 MCP 服务器脚手架
- 实现 MCP 协议的各种端点
- 处理认证、授权、请求验证
- 集成各种数据源和工具
- 生成 MCP 客户端 SDK

## 协议规范
MCP 服务器需要实现以下端点：
- GET /health: 健康检查
- GET /metadata: 服务器元数据
- POST /invoke: 调用工具
- POST /stream: 流式调用
- POST /cancel: 取消请求

## 使用示例
```go
// 创建 MCP 服务器
server := mcp.NewServer("my-server", "1.0.0")

// 注册工具
server.RegisterTool("tool-name", "Tool description", toolHandler)

// 启动服务器
server.Start(":8080")
```

## 配置选项
- 端口配置
- 认证方式（API Key, JWT）
- CORS 设置
- 速率限制
- 日志配置
