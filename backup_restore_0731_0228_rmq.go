// 代码生成时间: 2025-07-31 02:28:46
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/kataras/iris/v12"
)

// Backup 文件备份
func Backup(w http.ResponseWriter, r *http.Request) {
    // 获取文件路径
    path := r.URL.Query().Get("path")
    if path == "" {
        fmt.Fprintln(w, "请提供文件路径")
        return
    }

    // 检查文件是否存在
    if _, err := os.Stat(path); os.IsNotExist(err) {
        fmt.Fprintln(w, "文件不存在")
        return
    }

    // 读取文件内容
    fileContent, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Fprintln(w, "读取文件失败")
        return
    }

    // 计算文件的MD5
    md5sum := md5.Sum(fileContent)
    md5Str := hex.EncodeToString(md5sum[:])

    // 创建备份文件
    backupPath := path + ".bak"
    if err := ioutil.WriteFile(backupPath, fileContent, 0644); err != nil {
        fmt.Fprintln(w, "创建备份文件失败")
        return
    }

    // 返回备份文件信息
    fmt.Fprintf(w, "备份文件创建成功，路径：%s，MD5：%s
", backupPath, md5Str)
}

// Restore 文件恢复
func Restore(w http.ResponseWriter, r *http.Request) {
    // 获取备份文件路径
    backupPath := r.URL.Query().Get("backup_path")
    if backupPath == "" {
        fmt.Fprintln(w, "请提供备份文件路径")
        return
    }

    // 获取原文件路径
    originalPath := r.URL.Query().Get("original_path")
    if originalPath == "" {
        fmt.Fprintln(w, "请提供原文件路径")
        return
    }

    // 检查备份文件是否存在
    if _, err := os.Stat(backupPath); os.IsNotExist(err) {
        fmt.Fprintln(w, "备份文件不存在")
        return
    }

    // 读取备份文件内容
    backupContent, err := ioutil.ReadFile(backupPath)
    if err != nil {
        fmt.Fprintln(w, "读取备份文件失败")
        return
    }

    // 检查原文件是否存在
    if _, err := os.Stat(originalPath); os.IsNotExist(err) {
        fmt.Fprintln(w, "原文件不存在")
        return
    }

    // 恢复文件
    if err := ioutil.WriteFile(originalPath, backupContent, 0644); err != nil {
        fmt.Fprintln(w, "文件恢复失败")
        return
    }

    // 返回恢复结果
    fmt.Fprintf(w, "文件恢复成功，路径：%s
", originalPath)
}

func main() {
    app := iris.New()

    // 注册备份接口
    app.Post("/backup", Backup)

    // 注册恢复接口
    app.Post("/restore", Restore)

    // 启动服务器
    log.Fatal(app.Listen(":8080"))
}
