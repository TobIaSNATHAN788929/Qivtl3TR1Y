// 代码生成时间: 2025-09-22 19:00:49
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
)

// RequestData 是请求数据的结构体
type RequestData struct {
    Name string `json:"name"`
}

func main() {
    // 创建一个 Iris 应用程序
    app := iris.New()

    // 注册一个 GET 请求处理器，路径为 /hello
    app.Get("/hello", func(ctx iris.Context) {
        // 从请求的查询参数中获取 'name' 的值
        name := ctx.URLParam("name")
        // 如果 'name' 参数不存在，则返回一个错误信息
        if name == "" {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString(fmt.Sprintf("Error: 'name' parameter is required."))
        } else {
            // 向客户端返回一条欢迎信息
            ctx.WriteString(fmt.Sprintf("Hello, %s!", name))
        }
    })

    // 注册一个 POST 请求处理器，路径为 /submit
    app.Post("/submit", func(ctx iris.Context) {
        // 声明一个 RequestData 实例
        var data RequestData
        // 解析请求体中的 JSON 数据到 RequestData 实例
        if err := ctx.ReadJSON(&data); err != nil {
            // 如果解析失败，则返回错误信息
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString(fmt.Sprintf("Error: %s", err.Error()))
        } else {
            // 如果解析成功，返回提交的数据
            ctx.JSON(iris.StatusOK, iris.Map{
                "message": "Data received",
                "data": data,
            })
        }
    })

    // 启动 HTTP 服务器，监听端口 8080
    app.Listen(":8080")
}
