// 代码生成时间: 2025-08-20 19:07:59
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "log"
)

// ShoppingCart represents a shopping cart with items.
type ShoppingCart struct {
    Items map[string]int
}

// AddItem adds a new item to the shopping cart.
func (s *ShoppingCart) AddItem(item string, quantity int) error {
    if quantity <= 0 {
        return fmt.Errorf("quantity must be greater than 0")
    }
    // Add the item to the cart's item map.
    s.Items[item] += quantity
    return nil
}

// RemoveItem removes an item from the shopping cart.
func (s *ShoppingCart) RemoveItem(item string) error {
    if _, exists := s.Items[item]; !exists {
        return fmt.Errorf("item does not exist in the cart")
    }
    // Remove the item from the cart's item map.
    delete(s.Items, item)
    return nil
}

// GetTotal returns the total number of items in the cart.
func (s *ShoppingCart) GetTotal() int {
    total := 0
    for _, quantity := range s.Items {
        total += quantity
    }
    return total
}

func main() {
    app := iris.New()

    // Initialize the shopping cart.
    cart := ShoppingCart{Items: make(map[string]int)}

    // Define a route to add an item to the cart.
    app.Post("/cart/add", func(ctx iris.Context) {
        item := ctx.URLParam("item")
        quantityStr := ctx.URLParam("quantity")
        quantity, err := strconv.Atoi(quantityStr)
        if err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "Invalid quantity. Please enter an integer.",
            })
            return
        }
        if err := cart.AddItem(item, quantity); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.StatusCode(iris.StatusOK)
        ctx.JSON(iris.Map{
            "message": "Item added successfully.",
            "total": cart.GetTotal(),
        })
    })

    // Define a route to remove an item from the cart.
    app.Post("/cart/remove", func(ctx iris.Context) {
        item := ctx.URLParam("item")
        if err := cart.RemoveItem(item); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.StatusCode(iris.StatusOK)
        ctx.JSON(iris.Map{
            "message": "Item removed successfully.",
            "total": cart.GetTotal(),
        })
    })

    // Start the server.
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatal(err)
    }
}
