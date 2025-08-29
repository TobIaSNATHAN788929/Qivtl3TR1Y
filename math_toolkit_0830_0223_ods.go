// 代码生成时间: 2025-08-30 02:23:48
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "github.com/kataras/iris/v12"
)

// MathService 结构体，包含数学计算工具集功能
type MathService struct {
}

// Add 提供加法服务
func (s *MathService) Add(ctx iris.Context) {
    sum := ctx.URLParam("sum")
    result, err := calculateSum(sum)
    if err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
    } else {
        ctx.JSON(iris.Map{
            "result": result,
        })
    }
}

// calculateSum 计算加法结果
func calculateSum(sum string) (float64, error) {
    values := strings.Split(sum, "+")
    total := 0.0
    for _, val := range values {
        num, err := strconv.ParseFloat(val, 64)
        if err != nil {
            return 0, fmt.Errorf("invalid number: %s", val)
        }
        total += num
    }
    return total, nil
}

// Subtract 提供减法服务
func (s *MathService) Subtract(ctx iris.Context) {
    minuend := ctx.URLParam("minuend")
    subtrahend := ctx.URLParam("subtrahend")
    result, err := calculateDifference(minuend, subtrahend)
    if err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
    } else {
        ctx.JSON(iris.Map{
            "result": result,
        })
    }
}

// calculateDifference 计算减法结果
func calculateDifference(minuend, subtrahend string) (float64, error) {
    num1, err := strconv.ParseFloat(minuend, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid minuend: %s", minuend)
    }
    num2, err := strconv.ParseFloat(subtrahend, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid subtrahend: %s", subtrahend)
    }
    return num1 - num2, nil
}

// Multiply 提供乘法服务
func (s *MathService) Multiply(ctx iris.Context) {
    multiplicand := ctx.URLParam("multiplicand")
    multiplier := ctx.URLParam("multiplier")
    result, err := calculateProduct(multiplicand, multiplier)
    if err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
    } else {
        ctx.JSON(iris.Map{
            "result": result,
        })
    }
}

// calculateProduct 计算乘法结果
func calculateProduct(multiplicand, multiplier string) (float64, error) {
    num1, err := strconv.ParseFloat(multiplicand, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid multiplicand: %s", multiplicand)
    }
    num2, err := strconv.ParseFloat(multiplier, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid multiplier: %s", multiplier)
    }
    return num1 * num2, nil
}

// Divide 提供除法服务
func (s *MathService) Divide(ctx iris.Context) {
    dividend := ctx.URLParam("dividend")
    divisor := ctx.URLParam("divisor")
    result, err := calculateQuotient(dividend, divisor)
    if err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
    } else {
        ctx.JSON(iris.Map{
            "result": result,
        })
    }
}

// calculateQuotient 计算除法结果
func calculateQuotient(dividend, divisor string) (float64, error) {
    num1, err := strconv.ParseFloat(dividend, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid dividend: %s", dividend)
    }
    num2, err := strconv.ParseFloat(divisor, 64)
    if err != nil {
        return 0, fmt.Errorf("invalid divisor: %s", divisor)
    }
    if num2 == 0 {
        return 0, fmt.Errorf("divisor cannot be zero")
    }
    return num1 / num2, nil
}

// RandomNumber 提供生成随机数的服务
func (s *MathService) RandomNumber(ctx iris.Context) {
    number, err := generateRandomNumber()
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
    } else {
        ctx.JSON(iris.Map{
            "result": number,
        })
    }
}

// generateRandomNumber 生成随机数
func generateRandomNumber() (float64, error) {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64(), nil
}

func main() {
    app := iris.New()

    // 定义路由
    mathService := MathService{}
    app.Get("/add", mathService.Add)
    app.Get("/subtract", mathService.Subtract)
    app.Get("/multiply", mathService.Multiply)
    app.Get("/divide", mathService.Divide)
    app.Get("/random", mathService.RandomNumber)

    // 启动服务器
    app.Listen(":8080")
}
