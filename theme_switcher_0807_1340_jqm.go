// 代码生成时间: 2025-08-07 13:40:22
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/kataras/iris/v12" // Import the IRIS framework
)

// ThemeData represents the data structure for theme settings.
type ThemeData struct {
    Theme string `json:"theme"`
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html")) // Register the HTML engine

    // Define a route for the theme switcher.
    app.Get("/switch-theme", func(ctx iris.Context) {
        // Get the theme from the query string.
        theme := ctx.URLParam("theme")

        // Check for errors and set the theme in the session.
        if len(theme) == 0 {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.Writef("Error: No theme provided")
            return
        }

        // Set the theme in the session (as an example, using a cookie).
        ctx.SetCookieKV("theme", theme)

        // Redirect to the home page with the new theme.
        ctx.Redirect("/")
    })

    // Define a route for the home page that uses the theme.
    app.Get("/", func(ctx iris.Context) {
        // Retrieve the theme from the session (cookie).
        theme := ctx.GetCookie("theme")
        if theme == "" {
            theme = "default" // Default theme if no cookie is set.
        }

        // Render the home page template with the theme.
        err := ctx.View("home.html", ThemeData{Theme: theme})
        if err != nil {
            log.Printf("Error rendering home page: %v", err)
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.Writef("Error: Unable to render home page")
            return
        }
    })

    // Start the IRIS server.
    log.Fatal(app.Listen(":8080"))
}
