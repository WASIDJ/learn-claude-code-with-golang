---
name: code-review
description: 代码审查，检查代码质量、安全性、性能等问题
---

# 代码审查技能

## 功能说明
- 静态代码分析
- 检查代码规范和最佳实践
- 安全漏洞扫描
- 性能问题检测
- 代码复杂度分析
- 测试覆盖率检查
- 自动生成审查报告

## 审查维度
1. **代码质量**
   - 命名规范
   - 代码重复度
   - 函数/类复杂度
   - 注释完整性

2. **安全性**
   - SQL 注入风险
   - XSS 漏洞
   - 认证授权问题
   - 敏感信息泄露

3. **性能**
   - 慢查询
   - 内存泄漏
   - 并发问题
   - 资源未释放

4. **可维护性**
   - 架构合理性
   - 依赖管理
   - 测试覆盖
   - 文档完整性

## 使用示例
```go
// 审查 Go 代码
report, err := codeReview.ReviewGo("path/to/project")

// 审查 JavaScript 代码
report, err := codeReview.ReviewJS("path/to/project")

// 生成 HTML 报告
report.GenerateHTML("review-report.html")
```

## 支持语言
- Go
- JavaScript/TypeScript
- Python
- Java
- Rust
