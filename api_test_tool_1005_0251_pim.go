// 代码生成时间: 2025-10-05 02:51:21
package main

import (
    "fmt"
    "net/http"
    "github.com/kataras/iris/v12"
)

// APITestTool 结构体，用于存放测试工具的配置和数据
type APITestTool struct{}

// NewAPITestTool 函数，用于创建一个新的 APITestTool 实例
func NewAPITestTool() *APITestTool {
    return &APITestTool{}
}

// SetupRoutes 函数，用于设置 API 测试工具的路由
func (tool *APITestTool) SetupRoutes(app *iris.Application) {
    // 设置 GET 请求测试路由
    app.Get("/test", func(ctx iris.Context) {
        // 测试数据
        testResponse := map[string]interface{}{
            "message": "Hello, this is a test response!",
        }
        // 返回测试响应
        ctx.JSON(http.StatusOK, testResponse)
    })

    // 设置 POST 请求测试路由
    app.Post("/test", func(ctx iris.Context) {
        // 从请求体中读取数据
        var requestBody struct {
            Name string `json:"name"`
            Age  int    `json:"age"`
        }
        // 解析请求体
        if err := ctx.ReadJSON(&requestBody); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            return
        }
        // 创建响应数据
        testResponse := map[string]interface{}{
            "message": "Received POST request",
            "data": requestBody,
        }
        // 返回测试响应
        ctx.JSON(http.StatusOK, testResponse)
    })
}

func main() {
    // 创建 Iris 应用实例
    app := iris.New()

    // 创建 API 测试工具实例
    tool := NewAPITestTool()

    // 设置路由
    tool.SetupRoutes(app)

    // 启动服务器
    fmt.Println("API Test Tool is running at :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Println("Failed to start API Test Tool: ", err)
    }
}