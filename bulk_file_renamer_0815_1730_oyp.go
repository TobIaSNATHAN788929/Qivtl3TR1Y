// 代码生成时间: 2025-08-15 17:30:26
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// FileRenamer 结构体，用于批量重命名文件
type FileRenamer struct {
    // 旧文件名和新文件名的映射
    renameMap map[string]string
}

// NewFileRenamer 创建一个新的FileRenamer实例
func NewFileRenamer() *FileRenamer {
    return &FileRenamer{
        renameMap: make(map[string]string),
    }
}

// AddMapping 添加一个文件重命名映射
func (fr *FileRenamer) AddMapping(oldName, newName string) {
    fr.renameMap[oldName] = newName
}

// RenameFiles 执行批量文件重命名
func (fr *FileRenamer) RenameFiles() error {
    for oldName, newName := range fr.renameMap {
        if err := renameFile(oldName, newName); err != nil {
            return fmt.Errorf("error renaming file: %w", err)
        }
    }
    return nil
}

// renameFile 重命名单个文件
func renameFile(oldPath, newPath string) error {
    if _, err := os.Stat(oldPath); os.IsNotExist(err) {
        return fmt.Errorf("file %s does not exist", oldPath)
    }
    if _, err := os.Stat(newPath); err == nil {
        return fmt.Errorf("file %s already exists", newPath)
    }
    if err := os.Rename(oldPath, newPath); err != nil {
        return fmt.Errorf("failed to rename file %s to %s: %w", oldPath, newPath, err)
    }
    return nil
}

func main() {
    // 示例使用
    fr := NewFileRenamer()
    fr.AddMapping("oldfile1.txt", "newfile1.txt")
    fr.AddMapping("oldfile2.txt", "newfile2.txt")

    // 执行批量重命名
    if err := fr.RenameFiles(); err != nil {
        fmt.Printf("Bulk file renaming failed: %s
", err)
    } else {
        fmt.Println("Bulk file renaming completed successfully")
    }
}
