// 代码生成时间: 2025-08-23 01:58:33
package main

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// CartItem represents an item in the shopping cart.
type CartItem struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Quantity    int       `json:"quantity"`
    Price       float64   `json:"price"`
    CreatedDate time.Time `json:"createdDate"`
}

// Cart represents the shopping cart.
type Cart struct {
    Items map[string]CartItem `json:"items"`
}

// NewCart initializes a new shopping cart.
func NewCart() *Cart {
    return &Cart{
        Items: make(map[string]CartItem),
    }
}

// AddItem adds an item to the cart.
func (c *Cart) AddItem(item CartItem) {
    c.Items[item.ID] = item
}

// RemoveItem removes an item from the cart.
func (c *Cart) RemoveItem(itemID string) {
    delete(c.Items, itemID)
}

// ServeHTTP handles HTTP requests for the shopping cart.
func (c *Cart) ServeHTTP(ctx iris.Context) {
    // Get the cart data from the session.
    cartData, err := ctx.Session().GetString("shoppingCart")
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Error retrieving cart from session.")
        return
    }

    // Decode the cart data into a Cart object.
    var cart Cart
    if err := json.Unmarshal([]byte(cartData), &cart); err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Error decoding cart data.")
        return
    }

    // Handle the request.
    switch ctx.Method() {
    case http.MethodGet:
        // Return the current cart contents.
        ctx.JSON(cart)
    case http.MethodPost:
        // Add an item to the cart.
        var newItem CartItem
        if err := ctx.ReadJSON(&newItem); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString("Error reading new item data.")
            return
        }
        cart.AddItem(newItem)
        ctx.Session().SetString("shoppingCart", cart.String())
        ctx.JSON(cart)
    }
}

// String returns a JSON string representation of the cart.
func (c *Cart) String() string {
    bytes, _ := json.Marshal(c)
    return string(bytes)
}

func main() {
    // Initialize the IRIS application.
    app := iris.New()

    // Create a new shopping cart.
    cart := NewCart()
    // Set up the session middleware.
    app.Use(iris.NewSessionMiddleware(iris.Config{
        Cookie:      "ShoppingCartSession",
        AllowSubdomain: true,
    }))

    // Define the cart route.
    app.Handle("GET", "/cart", func(ctx iris.Context) {
        cart.ServeHTTP(ctx)
    })
    app.Handle("POST", "/cart", func(ctx iris.Context) {
        cart.ServeHTTP(ctx)
    })

    // Start the IRIS server.
    log.Fatal(app.Run(iris.Addr(":8080")))
}
