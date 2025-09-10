// 代码生成时间: 2025-09-11 04:21:46
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    "github.com/kataras/iris/v12/middleware/logger"
)

// 定义表单数据验证的错误类型
type ValidationError struct {
    Field string
    Error string
}

// FormValidator 结构体用于封装表单验证逻辑
type FormValidator struct {
    errors []ValidationError
}

// AddError 添加一个验证错误到验证器中
func (v *FormValidator) AddError(field, errMsg string) {
    v.errors = append(v.errors, ValidationError{Field: field, Error: errMsg})
}

// IsValid 检查是否有验证错误，如果没有返回true，否则返回false
func (v *FormValidator) IsValid() bool {
    return len(v.errors) == 0
}

// ValidateForm 验证表单数据
func (v *FormValidator) ValidateForm(ctx iris.Context, fields map[string]string) bool {
    // 示例：验证字段非空
    for field, value := range fields {
        if value == "" {
            v.AddError(field, "Field cannot be empty")
        }
    }
    // 可以添加更多的验证规则
    return v.IsValid()
}

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    // 定义表单数据验证器
    validator := &FormValidator{}

    // 设置表单提交路由
    app.Post("/form", func(ctx iris.Context) {
        fields := map[string]string{
            "username": ctx.FormValue("username"),
            "password": ctx.FormValue("password"),
        }
        if validator.ValidateForm(ctx, fields) {
            fmt.Println("Form is valid")
            ctx.JSON(iris.StatusOK, iris.Map{
                "message": "Form is valid",
            })
        } else {
            fmt.Println("Form is invalid")
            ctx.JSON(iris.StatusBadRequest, iris.Map{
                "errors": validator.errors,
            })
        }
    })

    // 启动服务器
    app.Listen(":8080")
}