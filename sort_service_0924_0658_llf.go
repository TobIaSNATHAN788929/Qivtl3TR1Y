// 代码生成时间: 2025-09-24 06:58:06
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/kataras/iris/v12" // Ensure you have this package installed
)

// SortAlgorithm defines the sorting algorithm to be used.
type SortAlgorithm string

const (
    // SelectionSort is the selection sort algorithm.
    SelectionSort SortAlgorithm = "selection"
    // BubbleSort is the bubble sort algorithm.
    BubbleSort SortAlgorithm = "bubble"
)

// SortService defines the interface for sorting services.
type SortService interface {
    Sort(numbers []int) ([]int, error)
}

// sortingService implements the SortService interface.
type sortingService struct {}

// NewSortingService creates a new sorting service instance.
func NewSortingService() SortService {
    return &sortingService{}
}

// Sort sorts the numbers using the specified algorithm.
func (s *sortingService) Sort(numbers []int, algorithm SortAlgorithm) ([]int, error) {
    switch algorithm {
    case SelectionSort:
        return selectionSort(numbers), nil
    case BubbleSort:
        return bubbleSort(numbers), nil
    default:
        return nil, fmt.Errorf("unsupported sorting algorithm: %s", algorithm)
    }
}

// selectionSort sorts numbers using the selection sort algorithm.
func selectionSort(numbers []int) []int {
    for i := 0; i < len(numbers); i++ {
        minIdx := i
        for j := i + 1; j < len(numbers); j++ {
            if numbers[j] < numbers[minIdx] {
                minIdx = j
            }
        }
        numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
    }
    return numbers
}

// bubbleSort sorts numbers using the bubble sort algorithm.
func bubbleSort(numbers []int) []int {
    n := len(numbers)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }
        }
    }
    return numbers
}

func main() {
    app := iris.New()
    
    // Seed the random number generator.
    rand.Seed(time.Now().UnixNano())
    
    // Example usage of sorting service.
    sortService := NewSortingService()
    exampleNumbers := generateRandomNumbers(10)
    sortedNumbers, err := sortService.Sort(exampleNumbers, SelectionSort)
    if err != nil {
        fmt.Printf("Error sorting numbers: %s
", err)
    } else {
        fmt.Printf("Sorted numbers: %+v
", sortedNumbers)
    }
    
    // Define routes and start the server.
    app.Get("/sort", func(ctx iris.Context) {
        // Generate a random slice of numbers to sort.
        numbers := generateRandomNumbers(10)
        // Sort the numbers using the default algorithm.
        sortedNumbers, err := sortService.Sort(numbers, SelectionSort)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Error sorting numbers")
        } else {
            ctx.JSON(iris.Map{
                "numbers": sortedNumbers,
            })
        }
    })
    
    // Start the server on port 8080.
    app.Listen(":8080")
}

// generateRandomNumbers generates a slice of random numbers with the given size.
func generateRandomNumbers(size int) []int {
    numbers := make([]int, size)
    for i := range numbers {
        numbers[i] = rand.Intn(100) // Random numbers between 0 and 99.
    }
    return numbers
}
