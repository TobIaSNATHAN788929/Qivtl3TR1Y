// 代码生成时间: 2025-09-01 04:46:29
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// ResponseLayoutHandler 定义一个响应式布局的处理函数
func ResponseLayoutHandler(ctx iris.Context) {
    // 检查是否为移动端访问
    userAgent := ctx.GetHeader("User-Agent")
    isMobile := false
    if strings.Contains(userAgent, "Mobi") || strings.Contains(userAgent, "Android") {
        isMobile = true
    }

    // 根据是否为移动端返回不同的布局模板
    if isMobile {
        ctx.View("mobile_layout.html")
    } else {
        ctx.View("desktop_layout.html")
    }
}

func main() {
    // 创建一个新的Iris应用实例
    app := iris.New()

    // 设置视图模板文件所在目录
    app.RegisterView(iris.HTML("./templates", ".html")).Reload(true)

    // 定义路由和处理函数
    app.Get("/layout", ResponseLayoutHandler)

    // 启动服务器
    fmt.Println("Server is running at :8080")
    app.Listen(":8080")
}
