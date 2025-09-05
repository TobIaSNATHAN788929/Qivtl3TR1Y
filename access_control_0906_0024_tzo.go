// 代码生成时间: 2025-09-06 00:24:04
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/jwt"
)

// Define a custom error type for unauthorized access
type UnauthorizedError struct{
# 扩展功能模块
    message string
# NOTE: 重要实现细节
}

// Error method satisfies the error interface
func (e *UnauthorizedError) Error() string {
    return e.message
}

// Define a Claims struct to hold our JWT claims
type Claims struct {
    jwt.Tokens
    UserID string `json:"user_id"`
}

func main() {
    app := iris.New()
    // Set up the JWT middleware
    authMiddleware := jwt.New(jwt.Config{
        SigningKey:   []byte("secret"),
        ContextKey:   "auth",
        ErrorContext: "error",
        Claims:       &Claims{},
    })

    // Define a route for protected content
    app.Get("/protected", func(ctx iris.Context) {
        // Retrieve the claims from the context
        claims, ok := ctx.Get("auth").(*Claims)
        if !ok || claims == nil {
            ctx.StatusCode(iris.StatusUnauthorized)
            return
        }

        // Access the UserID from the claims
        userID := claims.UserID
        ctx.Writef("Hello, %s!", userID)
    }, authMiddleware)

    // Define a route to handle the error context
    app.On(iris.StatusUnauthorized, func(ctx iris.Context) {
        err, ok := ctx.Values().Get("error").(error)
        if !ok {
            err = &UnauthorizedError{message: "Unauthorized access"}
        }
        ctx.JSON(iris.StatusBadRequest, iris.NewMap{
# TODO: 优化性能
            "error": err.Error(),
        })
    })

    // Start the server
    fmt.Println("Server is running on :8080")
    app.Listen(":8080")
}
