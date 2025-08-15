// 代码生成时间: 2025-08-16 01:13:18
package main

import (
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "strconv"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/logger"
    "github.com/kataras/iris/v12/middleware/recover"
)

// ImageResizer defines the struct for image resizing functionality.
type ImageResizer struct {
    // Width and Height specify the target dimensions of the resized image.
    Width, Height int
}

// NewImageResizer creates a new ImageResizer with the specified width and height.
func NewImageResizer(width, height int) *ImageResizer {
    return &ImageResizer{
        Width:  width,
        Height: height,
    }
}

// ResizeImage resizes an image to the specified dimensions.
func (r *ImageResizer) ResizeImage(img image.Image) image.Image {
    // Create a new image with the desired size.
    newImg := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))
    // Resize the image using the standard library's resizing algorithm.
    newImg = resize.Resize(uint(r.Width), uint(r.Height), img, resize.NearestNeighbor)
    return newImg
}

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    // Define the image resizer.
    resizer := NewImageResizer(800, 600)

    // Define the route for image resizing.
    app.Post("/resize", func(ctx iris.Context) {
        // Get the uploaded file.
        file, info, err := ctx.FormFile("image")
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Error retrieving the image file.")
            return
        }
        defer file.Close()

        // Check if the file is an image.
        if !isImageFile(info.Name) {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString("Unsupported file type.")
            return
        }

        // Open the image file.
        img, err := os.Open(file.Name())
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Error opening the image file.")
            return
        }
        defer img.Close()

        // Decode the image.
        imgFile, _, err := image.Decode(img)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Error decoding the image.")
            return
        }

        // Resize the image.
        resizedImg := resizer.ResizeImage(imgFile)

        // Save the resized image.
        output := "resized_image." + getExtension(info.Name)
        out, err := os.Create(output)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Error creating the resized image file.")
            return
        }
        defer out.Close()

        // Encode the resized image to the specified format and write to the output file.
        if strings.HasSuffix(info.Name, ".png") {
            png.Encode(out, resizedImg)
        } else {
            jpeg.Encode(out, resizedImg, &jpeg.Options{Quality: 100})
        }

        // Respond with the path to the resized image.
        ctx.JSON(iris.StatusOK, iris.Map{
            "resizedImagePath": output,
        })
    })

    // Start the server.
    app.Listen(":8080")
}

// isImageFile checks if the file is an image based on its extension.
func isImageFile(filename string) bool {
    return strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".jpeg") || strings.HasSuffix(filename, ".png")
}

// getExtension returns the file extension.
func getExtension(filename string) string {
    return filepath.Ext(filename)
}

// resize is a dependency that handles the resizing operation.
import _ "github.com/nfnt/resize"
