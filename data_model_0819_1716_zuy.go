// 代码生成时间: 2025-08-19 17:16:09
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// User 是我们的数据模型
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// NewUser 创建一个新的用户实例
func NewUser(id uint, username, email string) User {
    return User{
        ID:       id,
        Username: username,
        Email:    email,
    }
}

// GetUsers 返回一个用户列表
func GetUsers() []User {
    users := []User{
        NewUser(1, "John Doe", "john@example.com"),
        NewUser(2, "Jane Doe", "jane@example.com"),
    }
    return users
}

// CreateUser 创建一个新的用户
func CreateUser(ctx iris.Context) {
    var newUser User
    if err := ctx.ReadJSON(&newUser); err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{"error": "Failed to read JSON"})
        return
    }
    // 这里可以添加数据库逻辑来保存新用户
    // 例如: err := database.Save(&newUser)
    // if err != nil {
    //     ctx.StatusCode(iris.StatusInternalServerError)
    //     ctx.JSON(iris.Map{"error": "Failed to create user"})
    //     return
    // }
    ctx.JSON(newUser)
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))
    
    // 定义路由
    app.Get("/users", func(ctx iris.Context) {
        users := GetUsers()
        ctx.JSON(users)
    })
    
    app.Post("/users", CreateUser)
    
    // 启动服务
    app.Listen(":8080")
}
