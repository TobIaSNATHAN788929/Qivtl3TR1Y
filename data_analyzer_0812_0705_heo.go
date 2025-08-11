// 代码生成时间: 2025-08-12 07:05:52
 * Features:
 * - Error handling
 * - Comments and documentation
 * - Adherence to GoLang best practices
 * - Maintainability and extensibility
 */

package main

import (
    "fmt"
    "math"
    "net/http"
    "encoding/json"
    "github.com/kataras/iris/v12"
)

// Define a struct to hold basic statistical data
type StatisticalData struct {
    Count    int     `json:"count"`
    Sum      float64 `json:"sum"`
    Average  float64 `json:"average"`
    Median   float64 `json:"median"`
    Variance float64 `json:"variance"`
    StandardDeviation float64 `json:"std_dev"`
}

// CalculateStats calculates basic statistics for a slice of numbers
func CalculateStats(numbers []float64) (*StatisticalData, error) {
    length := len(numbers)
    if length == 0 {
        return nil, fmt.Errorf("no data to analyze")
    }

    sum := 0.0
    for _, num := range numbers {
        sum += num
    }

    avg := sum / float64(length)
    var median float64
    if length%2 == 0 {
        mid1 := numbers[length/2]
        mid2 := numbers[length/2 - 1]
        median = (mid1 + mid2) / 2
    } else {
        median = numbers[length/2]
    }

    var variance float64
    for _, num := range numbers {
        variance += math.Pow(num - avg, 2)
    }
    variance /= float64(length)

    stdDev := math.Sqrt(variance)

    return &StatisticalData{
        Count:    length,
        Sum:      sum,
        Average:  avg,
        Median:   median,
        Variance: variance,
        StandardDeviation: stdDev,
    }, nil
}

func main() {
    app := iris.New()

    // Define a route to handle POST requests with JSON data
    app.Post("/analyze", func(ctx iris.Context) {
        var numbers []float64
        // Unmarshal JSON body into the numbers slice
        if err := ctx.ReadJSON(&numbers); err != nil {
            ctx.JSON(http.StatusInternalServerError, iris.Map{"error": "failed to read JSON"})
            return
        }

        // Calculate statistics
        stats, err := CalculateStats(numbers)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, iris.Map{"error": err.Error()})
            return
        }

        // Send the results as a JSON response
        ctx.JSON(http.StatusOK, stats)
    })

    // Start the IRIS server on port 8080
    if err := app.Run(iris.Addr:"8080"); err != nil {
        fmt.Printf("Failed to start server: %s
", err)
    }
}
