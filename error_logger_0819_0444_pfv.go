// 代码生成时间: 2025-08-19 04:44:35
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/kataras/iris/v12"
)

// ErrorLogger 中间件，用于记录错误日志
type ErrorLogger struct {
    // 文件路径和名称
    filePath string
    // 日志输出
    out *log.Logger
}

// NewErrorLogger 创建一个新的 ErrorLogger 实例
func NewErrorLogger(filePath string) *ErrorLogger {
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    return &ErrorLogger{
        filePath: filePath,
        out:      log.New(file, "ERROR: ", log.LstdFlags),
    }
}

// Serve 定义中间件的 Serve 方法
func (l *ErrorLogger) Serve(ctx iris.Context) {
    ctx.Next()
    if err := ctx.GetErr(); err != nil {
        l.out.Printf("[%s] %s
", time.Now().UTC().Format(time.RFC3339), err)
    }
}

func main() {
    // 定义日志路径
    logPath := filepath.Join(os.TempDir(), "error_log.txt")
    // 创建 ErrorLogger 实例
    errorLogger := NewErrorLogger(logPath)

    app := iris.New()
    // 注册中间件
    app.Use(errorLogger.Serve)

    // 测试路径，将触发错误
    app.Get("/test-error", func(ctx iris.Context) {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.Writef("%s", "Internal Server Error")
    })

    // 启动服务器
    app.Listen(":8080")
}
