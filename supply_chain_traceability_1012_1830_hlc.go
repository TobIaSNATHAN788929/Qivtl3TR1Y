// 代码生成时间: 2025-10-12 18:30:44
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
)

// Product represents a product in the supply chain.
type Product struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    BatchNumber string `json:"batch_number"`
    Manufacturer string `json:"manufacturer"`
    Supplier string `json:"supplier"`
}

// ProductService handles business logic for products.
type ProductService struct {}

// CreateProduct creates a new product in the supply chain.
func (ps *ProductService) CreateProduct(ctx iris.Context, product Product) (Product, error) {
    // Add business logic for creating a product
    // For simplicity, we're just returning the input as the output
    return product, nil
}

// GetProduct retrieves a product by its ID.
func (ps *ProductService) GetProduct(ctx iris.Context, id string) (Product, error) {
    // Add business logic for retrieving a product
    // For simplicity, we're just returning a dummy product
    return Product{ID: id, Name: "Example Product", BatchNumber: "123456", Manufacturer: "Example Manufacturer", Supplier: "Example Supplier"}, nil
}

func main() {
    app := iris.New()
    productService := ProductService{}

    // Register API routes
    app.Post("/product", func(ctx iris.Context) {
        var product Product
        if err := ctx.ReadJSON(&product); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            return
        }
        createdProduct, err := productService.CreateProduct(ctx, product)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            return
        }
        ctx.JSON(createdProduct)
    })

    app.Get("/product/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        product, err := productService.GetProduct(ctx, id)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            return
        }
        ctx.JSON(product)
    })

    // Start the Iris server
    fmt.Println("Now server is running at :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Println(err)
    }
}
