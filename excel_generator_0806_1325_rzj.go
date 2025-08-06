// 代码生成时间: 2025-08-06 13:25:54
package main

import (
    "os"
    "log"
    "fmt"
    "github.com/xuri/excelize/v2" // 导入excelize库
    "github.com/kataras/iris/v12"
)

// ExcelGenerator 定义Excel生成器结构
type ExcelGenerator struct {
    excel *excelize.File
}

// NewExcelGenerator 创建一个新的Excel生成器实例
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{
        excel: excelize.NewFile(),
    }
}

// GenerateExcel 创建一个新的Excel文件并添加数据
func (e *ExcelGenerator) GenerateExcel(data [][]string) (string, error) {
    // 创建Excel文件
    e.excel = excelize.NewFile()

    // 设置Excel文件的默认工作表
    if err := e.excel.SetActiveSheet(e.excel.NewSheet(0)); err != nil {
        return "", err
    }

    // 添加数据到Excel
    for i, row := range data {
        for j, value := range row {
            if err := e.excel.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), value); err != nil {
                return "", err
            }
        }
    }

    // 保存Excel文件
    filePath := fmt.Sprintf("./out_%d.xlsx", len(data))
    if err := e.excel.SaveAs(filePath); err != nil {
        return "", err
    }

    return filePath, nil
}

func main() {
    app := iris.New()

    // 设置静态文件服务
    app.StaticServe("/static", "./static")

    // 设置Excel生成器路由
    app.Post("/excel", func(ctx iris.Context) {
        // 模拟数据
        data := [][]string{
            {"Header1", "Header2", "Header3", "Header4"},
            {"Data1", "Data2", "Data3", "Data4"},
            {"Data5", "Data6", "Data7", "Data8"},
        }

        // 创建Excel生成器实例
        generator := NewExcelGenerator()

        // 生成Excel文件
        filePath, err := generator.GenerateExcel(data)
        if err != nil {
            log.Printf("Error generating Excel file: %v", err)
            ctx.JSON(iris.StatusInternalServerError, iris.Map{"error": "Failed to generate Excel file"})
            return
        }

        // 下载Excel文件
        ctx.Download(filePath, "output.xlsx")
    })

    // 启动IRIS服务器
    if err := app.Listen(":8080"); err != nil {
        log.Fatalf("Error starting IRIS server: %v", err)
    }
}
