// 代码生成时间: 2025-09-23 07:11:03
package main

import (
    "os"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    "github.com/kataras/iris/v12/middleware/logger"
)

// ChartGenerator 用于生成交互式图表的HTTP处理器
type ChartGenerator struct{}

// GetHandler 处理GET请求，生成交互式图表的页面
func (c *ChartGenerator) GetHandler(ctx iris.Context) {
    ctx.Render("chart.html", iris.Map{"Title": "Interactive Chart Generator"})
}

// PostHandler 处理POST请求，接收图表数据并生成图表
func (c *ChartGenerator) PostHandler(ctx iris.Context) {
    var data struct{
        Labels []string `json:"labels"`
        Values []int    `json:"values"`
    }
    if err := ctx.ReadJSON(&data); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
        return
    }
    // 生成图表的逻辑（示例）
    // 这里可以添加图表生成的代码，例如使用第三方图表库
    // 例如: chart := chart.NewChart()
    // chart.AddSeries(data.Labels, data.Values)
    // ...
    ctx.JSON(iris.Map{
        "status": "success",
        "message": "Chart generated successfully",
        "data": data,
    })
}

func newApp() *iris.Application {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())
    
    // 设置静态文件目录
    app.HandleDir("/static", iris.Dir("./static"))
    
    // 路由设置
    app.Get("/chart", (&ChartGenerator{}).GetHandler)
    app.Post("/chart", (&ChartGenerator{}).PostHandler)
    return app
}

func main() {
    app := newApp()
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to run app: %s", err)
    }
}