// 代码生成时间: 2025-08-14 15:08:58
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "gopkg.in/go-playground/validator.v10"
)

// Form struct represents the data from the form
type Form struct {
    Name    string `validate:"required,min=2,max=100"`
    Email   string `validate:"required,email"`
    Age     int    `validate:"required,gte=1,lte=100"`
    Comment string `validate:"max=500"`
}

// ValidateForm is a function that validates the form data
func ValidateForm(ctx iris.Context) {
    // Create an instance of the form
    formData := Form{}

    // Bind the form data from the request
    if err := ctx.ReadForm(&formData); err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
        return
    }

    // Validate the form data
    validate := validator.New()
    result := validate.Struct(formData)
    if result.HasAny() {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": result.Error().String(),
        })
        return
    }

    // If validation passes, respond with success
    ctx.JSON(iris.Map{
        "message": "Form data is valid",
        "data": formData,
    })
}

func main() {
    app := iris.New()
    app.Logger().SetLevel("debug")

    // Register the form validation route
    app.Post("/form", ValidateForm)

    // Start the Iris server
    app.Listen(":8080")
}
