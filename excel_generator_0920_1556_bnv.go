// 代码生成时间: 2025-09-20 15:56:47
package main

import (
    "encoding/json"
    "excelize"
    "fmt"
    "log"
    "net/http"
    "path/filepath"
    "time"
)

// ExcelGenerator 结构体用于定义Excel生成器
type ExcelGenerator struct {
    // 用于存储Excel文件的路径
    FilePath string
    // 工作簿名称
    WorksheetName string
}

// NewExcelGenerator 创建ExcelGenerator实例
func NewExcelGenerator(filePath, worksheetName string) *ExcelGenerator {
    return &ExcelGenerator{
        FilePath:       filePath,
        WorksheetName: worksheetName,
    }
}

// GenerateExcel 生成Excel文件
func (eg *ExcelGenerator) GenerateExcel(data interface{}) error {
    // 创建Excel文件
    f := excelize.NewFile()
    // 设置工作簿名称
    index := f.NewSheet(eg.WorksheetName)
    // 创建表头
    titles := []string{"Column1", "Column2", "Column3"}
    f.SetSheetRow(eg.WorksheetName, "A1", &excelize.CellValue{Value: titles})
    
    // 将数据写入Excel
    dataBytes, err := json.Marshal(data)
    if err != nil {
        return err
    }
    var jsonData []map[string]interface{}
    if err := json.Unmarshal(dataBytes, &jsonData); err != nil {
        return err
    }
    for i, record := range jsonData {
        for j, value := range record {
            // 计算单元格位置
            cell := fmt.Sprintf("%c%d", 'A'+j+1, i+2) // 跳过表头
            f.SetCellValue(eg.WorksheetName, cell, fmt.Sprintf("%v", value))
        }
    }
    
    // 保存文件
    if err := f.SaveAs(eg.FilePath); err != nil {
        return err
    }
    
    return nil
}

// Handler 结构体用于定义处理Excel生成的handler
type Handler struct {
    ExcelGenerator *ExcelGenerator
}

// NewHandler 创建Handler实例
func NewHandler(generator *ExcelGenerator) *Handler {
    return &Handler{ExcelGenerator: generator}
}

// ExcelGenerationHandler 处理Excel生成请求
func (h *Handler) ExcelGenerationHandler(w http.ResponseWriter, r *http.Request) {
    // 仅允许POST请求
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // 解析请求体
    var requestData interface{}
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid Request", http.StatusBadRequest)
        log.Printf("Error parsing request data: %v", err)
        return
    }
    
    // 生成Excel文件
    if err := h.ExcelGenerator.GenerateExcel(requestData); err != nil {
        http.Error(w, "Failed to generate Excel", http.StatusInternalServerError)
        log.Printf("Error generating Excel: %v", err)
        return
    }
    
    // 返回成功的响应
    fmt.Fprintf(w, "Excel file generated successfully")
}

func main() {
    // 创建Excel生成器实例
    generator := NewExcelGenerator(
        filepath.Join(".", "generated_excel.xlsx"),
        "Sheet1",
    )
    
    // 创建Handler实例
    handler := NewHandler(generator)
    
    // 设置Iris路由
   iris.Get("/generate_excel", func(ctx iris.Context) {
        handler.ExcelGenerationHandler(ctx.ResponseWriter(), ctx.Request())
    })
    
    // 启动服务
    log.Println("Starting Excel generator service on :8080")
    iris.Listen(":8080")
}