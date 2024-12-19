# 插件式微信 AI 工具

## 项目简介

这个项目是一个插件式的微信 AI 工具，旨在为微信用户提供多种 AI 服务。目前支持的 AI 服务包括：

- **OpenAI**（GPT 系列）
- **Azure**（微软认知服务）
- **DeepSeek**（深度搜索引擎）
- **ZhiPu**（智谱）

该工具的设计理念是通过插件化的方式，便于未来的功能扩展。用户只需要获取对应的 API Token，填写到配置文件中，即可轻松启动并使用相应的 AI 服务。

## 特性

- 插件化设计：每个功能都作为独立的插件，易于扩展和维护。
- 支持多种 AI 引擎：目前支持 OpenAI、Azure、DeepSeek 和 ZhiPu 等主流 AI 平台。
- 简单配置：只需获取 API Token 并填写到配置文件，快速启动工具。
- 高效解答：提供多种 AI 解答服务，满足不同用户需求。

## 运行调试

### 1. 克隆仓库

首先，你需要克隆该项目到本地：

```bash
git clone https://github.com/Haif-07/wechat-ai.git
cd wechat-ai
```
### 2. 修改配置文件
  1. token
  2. 指定微信群
  3. 指定提示词

### 3. 修改配置文件
  ```bash
go run main.go
```

## 引用
  1. [openwechat](https://github.com/eatmoreapple/openwechat)
  2. [openai](https://github.com/ysicing/openai)
