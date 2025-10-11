// 代码生成时间: 2025-10-11 19:36:47
package main

import (
    "fmt"
    "net/url"
    "strings"

    "github.com/kataras/iris/v12"
)

// URLValidator 结构体提供一个方法来验证URL是否有效
type URLValidator struct{}

// ValidateURL 检查给定的URL是否有效
func (v *URLValidator) ValidateURL(ctx iris.Context, urlStr string) (bool, error) {
    u, err := url.ParseRequestURI(urlStr)
    if err != nil {
        return false, fmt.Errorf("invalid URL format: %w", err)
    }
    switch u.Scheme {
    case "http", "https", "ftp", "ftps":
        // 以下协议认为是有效的
    case "":
        // 如果协议为空，认为是无效的
        return false, nil
    default:
        // 其他协议认为无效
        return false, nil
    }
    return true, nil
}

func main() {
    app := iris.New()

    // 初始化URLValidator
    urlValidator := &URLValidator{}

    // 定义GET路由，用户可以通过这个路由提交URL进行检查
    app.Get("/check-url", func(ctx iris.Context) {
        // 获取用户提交的URL参数
        urlStr := ctx.URLParam("url")

        // 验证URL是否有效
        valid, err := urlValidator.ValidateURL(ctx, urlStr)
        if err != nil {
            // 处理错误
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": "Internal Server Error",
            })
            return
        }

        // 返回验证结果
        ctx.JSON(iris.Map{
            "valid": valid,
        })
    })

    // 启动服务器
    app.Listen(":8080")
}
