// 代码生成时间: 2025-09-01 11:47:03
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/sessions"
    "golang.org/x/crypto/bcrypt"
)

// User represents the user model
type User struct {
    Username string
    Password string
}

// AuthHandler is the handler for user authentication
func AuthHandler(ctx iris.Context) {
    // Retrieve username and password from the request
    username := ctx.URLParam("username")
    password := ctx.URLParam("password")

    // Check if the username and password are provided
    if username == "" || password == "" {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": "Username and password are required",
        })
        return
    }

    // Simulate user verification (in a real application, this would involve database lookup)
    expectedUsername := "admin"
    expectedPassword := "password" // This should be hashed in real applications

    if username != expectedUsername || password != expectedPassword {
        ctx.StatusCode(http.StatusUnauthorized)
        ctx.JSON(iris.Map{
            "error": "Invalid username or password",
        })
        return
    }

    // Hash the provided password for comparison
    hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatalf("Error hashing password: %v", err)
    }

    // Compare hashed password with the expected one
    if err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(expectedPassword)); err != nil {
        ctx.StatusCode(http.StatusUnauthorized)
        ctx.JSON(iris.Map{
            {
                "error": "Invalid username or password",
            },
        })
        return
    }

    // Set session for user
    sess := sessions.Get(ctx)
    sess.Set("username", username)
    sess.Expire(time.Hour) // Set session to expire after 1 hour

    ctx.JSON(iris.Map{
        "message": "User authenticated successfully",
    })
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))
    sessions.Configure(app, sessions.Config{
        Cookie:       "irissession",
        Expires:     time.Hour,
        AllowReclaim: true,
    })

    // Define routes
    app.Get("/login", AuthHandler)

    // Start the server
    app.Listen(":8080")
}
