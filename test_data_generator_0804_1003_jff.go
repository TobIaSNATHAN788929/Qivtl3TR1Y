// 代码生成时间: 2025-08-04 10:03:04
 * It includes error handling, documentation, and follows Golang best practices for maintainability and scalability.
 */

package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/kataras/iris/v12" // IRIS framework
)

// TestData represents the structure of the generated test data
type TestData struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Age     int    `json:"age"`
    Created string `json:"created"`
}

// NewTestData creates a new TestData instance with random values
func NewTestData(id int) TestData {
    rand.Seed(time.Now().UnixNano())
    name := fmt.Sprintf("TestUser%d", id)
    email := fmt.Sprintf("%s@example.com", name)
    age := rand.Intn(100) // Random age between 0 and 99
    created := time.Now().Format(time.RFC3339)
    return TestData{
        ID:      id,
        Name:    name,
        Email:   email,
        Age:     age,
        Created: created,
    }
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    // Generate and display test data
    app.Get("/test-data", func(ctx iris.Context) {
        for i := 1; i <= 10; i++ { // Generate 10 test data entries
            testData := NewTestData(i)
            ctx.JSON(iris.StatusOK, testData)
        }
    })

    // Error handling middleware
    app.Use(func(ctx iris.Context) {
        error := ctx.Values().GetString("err")
        if error != "" {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{