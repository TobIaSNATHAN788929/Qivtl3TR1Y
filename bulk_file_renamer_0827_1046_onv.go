// 代码生成时间: 2025-08-27 10:46:44
// bulk_file_renamer.go 是一个使用GOLANG和IRIS框架实现的批量文件重命名工具。

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// RenamePattern 定义重命名的规则
type RenamePattern struct {
    Search string // 要搜索的字符串
    Replace string // 要替换的字符串
}

// renameFiles 批量重命名指定目录下的所有文件
func renameFiles(dir string, patterns []RenamePattern) error {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        filePath := filepath.Join(dir, file.Name())
        newFileName := file.Name()
        for _, pattern := range patterns {
            newFileName = strings.ReplaceAll(newFileName, pattern.Search, pattern.Replace)
        }

        if newFileName != file.Name() {
            newFilePath := filepath.Join(dir, newFileName)
            if err := os.Rename(filePath, newFilePath); err != nil {
                return fmt.Errorf("failed to rename file %s to %s: %w", filePath, newFilePath, err)
            }
            fmt.Printf("Renamed %s to %s
", filePath, newFilePath)
        }
    }

    return nil
}

func main() {
    // 定义重命名规则
    patterns := []RenamePattern{
        {Search: "old", Replace: "new"},
        // 可以添加更多的重命名规则
    }

    // 指定需要重命名文件的目录
    directory := "./files"

    if err := renameFiles(directory, patterns); err != nil {
        log.Fatalf("Error renaming files: %s", err)
    }
}
