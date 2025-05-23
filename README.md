# Ginit - Go Project Scaffolding Tool

Ginit 是一个强大的 Go 项目脚手架工具，用于快速创建和初始化各种类型的 Go 项目。它提供了简单易用的命令行界面，帮助开发者快速搭建标准的 Go 项目结构。

## 声明

本项目代码借鉴自 [go-nunu/nunu](https://github.com/go-nunu/nunu) 项目，仅用于个人学习使用。感谢原作者的贡献！


## 功能特点

- 支持多种项目模板（Advanced、Admin、Basic、Chat）
- 提供交互式命令行界面
- 自动配置 Go 模块
- 自动安装依赖
- 支持自定义模板仓库

## 安装

确保你的系统已安装 Go 1.16 或更高版本，然后运行：

```bash
go install github.com/liqianbro/Ginit@latest
```

## 使用方法

1. 基本使用：
```bash
ginit new
```

2. 直接指定项目名称：
```bash
ginit new my-project
```

3. 使用自定义模板仓库：
```bash
ginit new -r https://github.com/liqianbro/template-repo
```

## 项目模板

Ginit 提供以下项目模板：

- **Advanced**: 包含丰富的功能，如数据库、JWT、定时任务、数据迁移、测试等
- **Admin**: 快速搭建后台管理系统的模板
- **Basic**: 基础项目结构
- **Chat**: 简单的聊天室项目，包含 WebSocket/TCP 功能

## 命令行参数

- `-r, --repo-url`: 指定自定义模板仓库地址
- `-h, --help`: 显示帮助信息

## 使用示例

1. 创建新项目：
```bash
ginit new
# 将启动交互式命令行界面，引导你选择项目模板和输入项目名称
```

2. 直接创建指定名称的项目：
```bash
ginit new my-project
# 将创建一个名为 my-project 的 Go 项目
```

3. 使用自定义模板：
```bash
ginit new -r https://github.com/username/custom-template
# 将使用指定的模板仓库创建项目
```

## 项目初始化流程

1. 选择项目模板
2. 克隆模板仓库
3. 替换包名和模块名
4. 执行 `go mod tidy`
5. 安装 wire 工具
6. 清理 Git 历史

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
