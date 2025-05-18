# logger-adapter-go

💡 一个支持多日志库适配、支持文件大小/时间分割、上下文绑定、线程安全的 Go 日志中间层库。

## ✅ 特性

- 支持日志库适配：`slog`, `zap`, `logrus`，或自定义日志库
- 日志切割：
    - ✅ 按文件大小（内置 `lumberjack`）
    - ✅ 按时间周期（集成 `rotatelogs`）
- 线程安全的全局日志器管理
- 上下文绑定日志器，适用于高并发/请求隔离
- 丰富示例、测试用例与基准测试

## 🔧 安装

```bash
go get github.com/lk2023060901/logger-adapter-go

## 开始

Examples 目录下涵盖了快速开始的示例代码和高级教程。