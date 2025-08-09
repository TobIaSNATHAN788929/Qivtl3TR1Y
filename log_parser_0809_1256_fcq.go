// 代码生成时间: 2025-08-09 12:56:00
package main

import (
    "fmt"
    "os"
    "strings"
    "time"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/logger"
    "gopkg.in/natefinch/lumberjack.v2"
)

// LogEntry represents a single log entry with its fields
type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
}

// ParseLogEntry parses a log line into a LogEntry struct
func ParseLogEntry(line string) (*LogEntry, error) {
    parts := strings.SplitN(line, " ", 3)
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid log format")
    }

    timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
    if err != nil {
        return nil, err
    }

    level := parts[2]
    message := strings.Join(parts[3:], " ")

    return &LogEntry{
        Timestamp: timestamp,
        Level:     level,
        Message:   message,
    }, nil
}

// LogParserService is a service that handles log file parsing
type LogParserService struct {
    LogFilePath string
}

// NewLogParserService creates a new LogParserService
func NewLogParserService(logFilePath string) *LogParserService {
    return &LogParserService{
        LogFilePath: logFilePath,
    }
}

// ParseLogFile parses the entire log file and returns a slice of LogEntry
func (s *LogParserService) ParseLogFile() ([]LogEntry, error) {
    file, err := os.Open(s.LogFilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var entries []LogEntry
    for scanner.Scan() {
        line := scanner.Text()
        entry, err := ParseLogEntry(line)
        if err != nil {
            // You can choose to skip the line, log the error, or handle it as you see fit
            continue
        }
        entries = append(entries, *entry)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return entries, nil
}

func main() {
    app := iris.New()
    logFilePath := "path/to/your/logfile.log"
    parserService := NewLogParserService(logFilePath)

    // Setup a logger to capture request logs
    app.Use(logger.New())
    // Setup a file-based logger to capture application logs
    logWriter := lumberjack.Logger{
        Filename:   "./logfile.log",
        MaxSize:    500, // in MB
        MaxBackups: 3,
        MaxAge:     28, // in days
    }
    app.Logger().SetOutput(&logWriter)

    // Endpoint to parse the log file and return the results
    app.Get("/parse", func(ctx iris.Context) {
        entries, err := parserService.ParseLogFile()
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(entries)
    })

    // Start the server
    app.Listen(":8080")
}
