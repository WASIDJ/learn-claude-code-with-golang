# 技能目录规范

## 目录结构
```
.skills/
├── skill-name-1/
│   └── SKILL.md      # 技能定义文件
├── skill-name-2/
│   └── SKILL.md
└── README.md
```

## SKILL.md 格式规范
每个技能目录下必须包含 `SKILL.md` 文件，格式如下：

```markdown
---
name: 技能名称       # 技能唯一标识，用于调用
description: 技能描述 # 简短描述技能功能
---

# 技能标题

## 功能说明
详细描述技能的功能和用途。

## 使用示例
提供使用示例代码或命令。

## 其他章节
根据需要添加其他相关章节。
```

## 现有技能
- `pdf`: 处理 PDF 文件，包括读取、解析、转换等操作
- `mcp-builder`: 构建 MCP (Model Context Protocol) 服务器
- `code-review`: 代码审查，检查代码质量、安全性、性能等问题
