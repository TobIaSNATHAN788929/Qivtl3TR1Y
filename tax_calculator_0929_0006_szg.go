// 代码生成时间: 2025-09-29 00:06:22
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// TaxRate 表示税率
type TaxRate struct {
    Rate float64 `json:"rate"`
}

// TaxCalculator 结构体，用于税务计算
type TaxCalculator struct {
    taxRate TaxRate
}

// Calculate 方法计算税金
func (tc *TaxCalculator) Calculate(income float64) (float64, error) {
    // 检查税率是否有效
    if tc.taxRate.Rate <= 0 || tc.taxRate.Rate > 1 {
        return 0, fmt.Errorf("invalid tax rate: %f", tc.taxRate.Rate)
    }
    // 计算税金
    tax := income * tc.taxRate.Rate
    return tax, nil
}

// NewTaxCalculator 创建一个新的税务计算器
func NewTaxCalculator(rate float64) *TaxCalculator {
    return &TaxCalculator{
        taxRate: TaxRate{Rate: rate},
    }
}

func main() {
    // 设置税率
    taxRate := 0.22 // 22%
    calculator := NewTaxCalculator(taxRate)

    // 初始化 Iris
    app := iris.New()

    // 定义路由和处理函数
    app.Post("/calculate", func(ctx iris.Context) {
        var request struct {
            Income float64 `json:"income"`
        }
        if err := ctx.ReadJSON(&request); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "invalid request",
            })
            return
        }

        tax, err := calculator.Calculate(request.Income)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(iris.Map{
            "income": request.Income,
            "tax": tax,
        })
    })

    // 启动服务器
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Printf("Server failed to start: %s
", err)
    }
}