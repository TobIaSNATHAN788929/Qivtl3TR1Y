// 代码生成时间: 2025-08-28 00:16:47
package main

import (
    "fmt"
    "os"
    "runtime"
    "strings"
    "time"

    "github.com/kardianos/osext"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
    "github.com/kataras/iris/v12"
)

// SystemPerformanceMonitor 结构体封装系统性能监控数据
type SystemPerformanceMonitor struct {
    CPUUsage    float64
    MemUsage    float64
    DiskUsage  float64
    NetworkUsage float64
}

// GetSystemPerformance 获取系统性能数据
func GetSystemPerformance() (*SystemPerformanceMonitor, error) {
    var monitor SystemPerformanceMonitor

    // 获取CPU使用率
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return nil, err
    }
    monitor.CPUUsage = cpuPercent[0]

    // 获取内存使用率
    memStat, err := mem.VirtualMemory()
    if err != nil {
        return nil, err
    }
    monitor.MemUsage = memStat.UsedPercent

    // 获取磁盘使用率
    diskStat, err := disk.Usage("/")
    if err != nil {
        return nil, err
    }
    monitor.DiskUsage = diskStat.UsedPercent

    // 获取网络使用率
    netIOStat, err := net.IOCounters()
    if err != nil {
        return nil, err
    }
    monitor.NetworkUsage = float64(netIOStat.BytesSent+netIOStat.BytesRecv) / float64(netIOStat.BytesRecv)

    return &monitor, nil
}

// StartServer 启动性能监控服务器
func StartServer() {
    app := iris.New()
    app.Get("/system-performance", func(ctx iris.Context) {
        monitor, err := GetSystemPerformance()
        if err != nil {
            ctx.JSON(iris.StatusInternalServerError, iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(iris.StatusOK, iris.Map{
            "cpu_usage":    monitor.CPUUsage,
            "mem_usage":    monitor.MemUsage,
            "disk_usage":   monitor.DiskUsage,
            "network_usage": monitor.NetworkUsage,
        })
    })

    app.Listen(":8080")
}

// main 函数启动性能监控服务
func main() {
    // 获取程序路径
    appPath, err := osext.ExecutableFolder()
    if err != nil {
        fmt.Println("Failed to get executable folder: " + err.Error())
        return
    }
    fmt.Printf("Application Path: %s
", appPath)

    // 打印Go版本和环境变量
    fmt.Printf("Go Version: %s
", runtime.Version())
    fmt.Printf("GOPATH: %s
", os.Getenv("GOPATH"))
    fmt.Printf("GOROOT: %s
", os.Getenv("GOROOT"))

    // 启动服务器
    StartServer()
}
