---
name: pdf
description: 处理 PDF 文件，包括读取、解析、转换等操作
---

# PDF 处理技能

## 功能说明
- 读取 PDF 文件内容
- 提取 PDF 中的文本、图片、表格
- 转换 PDF 为其他格式（如文本、Markdown、Word）
- 合并、拆分 PDF 文件
- 添加水印、加密解密 PDF

## 使用示例
```go
// 读取 PDF 内容
content, err := pdf.Read("example.pdf")

// 提取 PDF 中的表格
tables, err := pdf.ExtractTables("example.pdf")

// 转换 PDF 为 Markdown
markdown, err := pdf.ToMarkdown("example.pdf")
```

## 依赖工具
- pdftotext (poppler-utils)
- pdfgrep
- qpdf
