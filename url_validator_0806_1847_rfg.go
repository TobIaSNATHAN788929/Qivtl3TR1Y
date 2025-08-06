// 代码生成时间: 2025-08-06 18:47:17
 * Usage:
 *  - Start the server and visit the route /validate?url=<your_url> to check if the URL is valid.
 */

package main

import (
    "net/http"
    "net/url"
    
    "github.com/kataras/iris/v12" // Import the IRIS framework
)

// ValidateURL checks if the provided URL is valid.
func ValidateURL(ctx iris.Context) {
    // Get the URL parameter from the query string.
    urlStr := ctx.URLParam("url")
    
    // Check if the URL parameter is empty.
    if urlStr == "" {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.WriteString("URL parameter is required")
        return
    }
    
    // Parse the URL to check its validity.
    parsedURL, err := url.ParseRequestURI(urlStr)
    if err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.WriteString("Invalid URL format")
        return
    }
    
    // Check if the scheme is valid (http, https).
    if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.WriteString("URL must use http or https scheme")
        return
    }
    
    // If all checks pass, return a success response.
    ctx.StatusCode(http.StatusOK)
    ctx.WriteString("URL is valid")
}

func main() {
    app := iris.New()
    
    // Define the route for URL validation.
    app.Get("/validate", ValidateURL)
    
    // Start the IRIS server.
    app.Listen(":8080")
}