// 代码生成时间: 2025-08-17 08:10:42
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "io/ioutil"
    "net/http"
)

// ScrapeContent scrapes content from a given URL.
// It returns the content as a string or an error if any occurs.
func ScrapeContent(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    // Read the body of the response
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    
    return string(content), nil
}

// StartWebScraper sets up the IRIS web framework and starts the server.
func StartWebScraper() {
    iris.Use(func(ctx iris.Context) {
        ctx.StatusCode(iris.StatusOK)
    })

    // Endpoint to scrape content from a given URL
    iris.Get("/scrape", func(ctx iris.Context) {
        url := ctx.URLParam("url")
        if url == "" {
            ctx.JSON(iris.Map{
                "error": "URL parameter is required",
            })
            return
        }

        content, err := ScrapeContent(url)
        if err != nil {
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        // Return the scraped content as a JSON response
        ctx.JSON(iris.Map{
            "content": content,
        })
    })

    // Start the iris server on port 8080
    iris.Listen(":8080")
}

func main() {
    StartWebScraper()
}
