// 代码生成时间: 2025-09-19 12:29:14
package main

import (
    "fmt"
    "testing"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/httptest"
)

// TestSuite provides a basic test suite for Iris applications.
// It setups up a test HTTP server and runs tests against it.
type TestSuite struct {
    server *iris.Application
    client *httptest.App
}

// SetupSuite initializes the test suite by creating an Iris application and a test client.
func (ts *TestSuite) SetupSuite(t *testing.T) {
    ts.server = iris.New()
    ts.client = httptest.New(ts.server, t)
}

// TearDownSuite tears down the test suite by shutting down the Iris server.
func (ts *TestSuite) TearDownSuite(t *testing.T) {
    ts.server.Shutdown(context.Background())
}

// TestExample tests a simple GET request to the /example endpoint.
func TestExample(t *testing.T) {
    // Setup the test suite.
    ts := new(TestSuite)
    ts.SetupSuite(t)
    defer ts.TearDownSuite(t)

    // Register a test route.
    ts.server.Get("/example", func(ctx iris.Context) {
        ctx.WriteString("Hello, World!")
    })

    // Perform the GET request.
    e := ts.client.GET("/example").Expect()
    e.Status(httptest.StatusOK).Body().Equal("Hello, World!")
}

// main function to run the tests.
func main() {
    testing.Main(func(mts *testing.M) {
        mts.Run()
    }, nil)
}