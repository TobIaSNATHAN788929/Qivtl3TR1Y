// 代码生成时间: 2025-09-24 14:05:39
 * integration_test_tool.go - This file contains a simple integration test tool using the IRIS framework in Go.
 */

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/httptest"
)

// TestHandler is a simple handler function for testing.
func TestHandler(ctx iris.Context) {
	ctx.WriteString("Hello, World!")
}

// TestSuite is a struct that holds the Iris application for testing.
type TestSuite struct {
	App *iris.Application
}

// NewTestSuite creates a new TestSuite.
func NewTestSuite() *TestSuite {
	app := iris.New()
	// Register TestHandler at the root path for demonstration purposes.
	app.Get("/", TestHandler)
	return &TestSuite{App: app}
}

// RunTests runs the integration tests.
func (ts *TestSuite) RunTests() {
	t := httptest.New(t, ts.App)
	defer t.Close()

	// Test home page.
	t.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("Hello, World!")
}

func main() {
	suite := NewTestSuite()
	
	// Set the current directory as the working directory for the Iris application.
	if err := os.Chdir(filepath.Dir(os.Args[0])); err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	// Run the Iris application in production mode for real-world usage.
	fmt.Println("Starting Iris application...")
	if err := suite.App.Run(iris.Addr(":8080")); err != nil {
		log.Fatalf("Failed to start Iris application: %v", err)
	}

	// Run the integration tests.
	fmt.Println("Running integration tests...")
	suite.RunTests()
}
