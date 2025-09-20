// 代码生成时间: 2025-09-21 04:46:12
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/go-iris/iris/v12"
    "github.com/go-iris/iris/v12/adaptors/httprouter"
    "github.com/go-iris/iris/v12/middleware/logger"
    "github.com/go-iris/iris/v12/middleware/recover"
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql" // 引入MySQL驱动
)

// 初始化数据库连接
var engine *xorm.Engine

// 初始化iris应用
func main() {
    // 设置数据库连接字符串
    dsn := "username:password@tcp(host:port)/dbname?charset=utf8"
    var err error
    engine, err = xorm.NewEngine("mysql", dsn)
    if err != nil {
        log.Fatalf("数据库连接失败: %v", err)
    }

    // 创建iris应用
    app := iris.New()
    app.Adapt(adaptors.NewHttpRouter())
    app.Use(recover.New())
    app.Use(logger.New())

    // 注册HTTP路由
    registerRoutes(app)

    // 启动iris应用
    app.Listen(":8080")
}

// 注册路由和处理函数
func registerRoutes(app *iris.Application) {
    // 假设有一个用户表User，包含字段ID和Name
    type User struct {
        ID   int64
        Name string `xorm:"'name'"` // 使用反引号防止SQL注入
    }

    // 获取用户信息的路由
    app.Get("/user", func(ctx iris.Context) {
        id := ctx.URLParam("id")
        // 使用参数化查询防止SQL注入
        user := new(User)
        has, err := engine.Where("id = ?", id).Get(user)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("数据库查询错误")
            return
        }
        if !has {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.WriteString("用户不存在")
            return
        }
        ctx.JSON(user)
    })
}
