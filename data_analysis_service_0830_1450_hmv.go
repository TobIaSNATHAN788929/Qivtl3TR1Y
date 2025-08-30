// 代码生成时间: 2025-08-30 14:50:31
package main

import (
    "fmt"
    "math"
    "time"
)

// DataPoint represents a single data point with a timestamp and a value.
type DataPoint struct {
    Timestamp time.Time
    Value     float64
}

// DataAnalysisService is the main struct that holds the data and provides methods for analysis.
type DataAnalysisService struct {
    Data []DataPoint
}

// NewDataAnalysisService creates a new instance of DataAnalysisService.
func NewDataAnalysisService() *DataAnalysisService {
    return &DataAnalysisService{
        Data: make([]DataPoint, 0),
    }
}

// AddDataPoint adds a new data point to the service.
func (das *DataAnalysisService) AddDataPoint(timestamp time.Time, value float64) {
    das.Data = append(das.Data, DataPoint{Timestamp: timestamp, Value: value})
}

// CalculateMean calculates the mean of all data points.
func (das *DataAnalysisService) CalculateMean() (float64, error) {
    if len(das.Data) == 0 {
        return 0, fmt.Errorf("no data points available")
    }
    sum := 0.0
    for _, dp := range das.Data {
        sum += dp.Value
    }
    return sum / float64(len(das.Data)), nil
}

// CalculateStandardDeviation calculates the standard deviation of all data points.
func (das *DataAnalysisService) CalculateStandardDeviation() (float64, error) {
    if len(das.Data) == 0 {
        return 0, fmt.Errorf("no data points available")
    }
    mean, err := das.CalculateMean()
    if err != nil {
        return 0, err
    }
    variance := 0.0
    for _, dp := range das.Data {
        variance += math.Pow(dp.Value-mean, 2)
    }
    return math.Sqrt(variance / float64(len(das.Data)-1)), nil
}

func main() {
    // Create a new data analysis service.
    das := NewDataAnalysisService()
    
    // Add some sample data points.
    das.AddDataPoint(time.Now(), 10.5)
    das.AddDataPoint(time.Now(), 20.3)
    das.AddDataPoint(time.Now(), 15.2)
    
    // Calculate and print the mean.
    mean, err := das.CalculateMean()
    if err != nil {
        fmt.Println("Error calculating mean: ", err)
    } else {
        fmt.Printf("Mean: %.2f
", mean)
    }
    
    // Calculate and print the standard deviation.
    stdDev, err := das.CalculateStandardDeviation()
    if err != nil {
        fmt.Println("Error calculating standard deviation: ", err)
    } else {
        fmt.Printf("Standard Deviation: %.2f
", stdDev)
    }
}
