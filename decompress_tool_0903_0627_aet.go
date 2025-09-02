// 代码生成时间: 2025-09-03 06:27:57
package main

import (
    "archive/zip"
    "fmt"
    "io"
    "io/fs"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/kataras/iris/v12"
)

// main is the entry point of the application.
func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    // Define the route for uploading and decompressing files.
    app.Post("/decompress", func(ctx iris.Context) {
        // Check if the file is present in the request form.
        formFile, _, err := ctx.FormFile("file")
        if err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.WriteString("No file selected")
            return
        }
        defer formFile.Close()

        // Save the uploaded file to a temporary location.
        tempFile, err := os.CreateTemp(os.TempDir(), "upload-*.zip")
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.WriteString("Error creating temporary file")
            return
        }
        defer tempFile.Close()
        defer os.Remove(tempFile.Name())

        // Copy the uploaded file to the temporary location.
        _, err = io.Copy(tempFile, formFile)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.WriteString("Error saving the file")
            return
        }

        // Decompress the file to the current directory.
        decompressedFiles, err := decompressZipFile(tempFile.Name())
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.WriteString("Error decompressing the file")
            return
        }

        // Send the list of decompressed files back to the client.
        ctx.JSON(iris.Map{
            "message": "Decompressed successfully",
            "files": decompressedFiles,
        })
    })

    // Start the IRIS web server.
    app.Listen(":8080")
}

// decompressZipFile decompresses a ZIP file to the current directory.
func decompressZipFile(zipFilePath string) ([]string, error) {
    var decompressedFiles []string
    reader, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return nil, err
    }
    defer reader.Close()

    for _, file := range reader.File {
        // Skip directories.
        if file.FileInfo().IsDir() {
            continue
        }

        // Create the directory structure.
        if err := os.MkdirAll(file.Name[:strings.LastIndex(file.Name, "/")], os.ModePerm); err != nil {
            return nil, err
        }

        // Create the file.
        outFile, err := os.Create(file.Name)
        if err != nil {
            return nil, err
        }
        defer outFile.Close()

        // Copy the contents of the ZIP file to the new file.
        fileReader, err := file.Open()
        if err != nil {
            return nil, err
        }
        defer fileReader.Close()
        _, err = io.Copy(outFile, fileReader)
        if err != nil {
            return nil, err
        }

        // Save the name of the decompressed file.
        decompressedFiles = append(decompressedFiles, file.Name)
    }
    return decompressedFiles, nil
}
