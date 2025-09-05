// 代码生成时间: 2025-09-05 19:34:37
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/logger"
    "github.com/kataras/iris/v12/middleware/recover"
)

// Component represents a UI component
type Component struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

// NewComponent creates a new UI component
func NewComponent(name, version string) *Component {
    return &Component{
        Name:    name,
        Version: version,
    }
}

func main() {
    // Initialize Iris
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    // Define routes for UI components
    app.Get("/components/:name", func(ctx iris.Context) {
        name := ctx.Param("name")

        // Simulate a database lookup for the component
        component := NewComponent(name, "1.0.0")
        if component.Name == "" {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.Map{
                "error": "Component not found",
            })
            return
        }

        // Return the component details
        ctx.JSON(component)
    })

    // Start the Iris server
    app.Listen(":8080")
}
