// 代码生成时间: 2025-09-09 18:38:19
package main

import (
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

// FolderOrganizer is a struct that holds the source directory path.
type FolderOrganizer struct {
    Source string
}

// NewFolderOrganizer creates a new instance of FolderOrganizer with the given source directory.
func NewFolderOrganizer(source string) *FolderOrganizer {
    return &FolderOrganizer{Source: source}
}

// Organize sorts files in the source directory by name and moves them into subdirectories.
func (f *FolderOrganizer) Organize() error {
    // Read the directory entries.
    entries, err := os.ReadDir(f.Source)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    // Create a map to hold the files grouped by the first letter of their names.
    filesByFirstLetter := make(map[rune][]fs.DirEntry)

    // Iterate over the directory entries and group them by the first letter.
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        firstLetter := rune(strings.ToUpper(entry.Name())[0])
        filesByFirstLetter[firstLetter] = append(filesByFirstLetter[firstLetter], entry)
    }

    // Sort the keys of the map.
    keys := make([]rune, 0, len(filesByFirstLetter))
    for k := range filesByFirstLetter {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

    // Create subdirectories and move files into them.
    for _, firstLetter := range keys {
        subdirPath := filepath.Join(f.Source, string(firstLetter))
        if err := os.MkdirAll(subdirPath, 0755); err != nil {
            return fmt.Errorf("failed to create subdirectory: %w", err)
        }
        for _, file := range filesByFirstLetter[firstLetter] {
            srcPath := filepath.Join(f.Source, file.Name())
            dstPath := filepath.Join(subdirPath, file.Name())
            if err := os.Rename(srcPath, dstPath); err != nil {
                return fmt.Errorf("failed to move file: %w", err)
            }
        }
    }

    return nil
}

func main() {
    // Define the source directory.
    sourceDir := "./source"

    // Create a new FolderOrganizer instance.
    organizer := NewFolderOrganizer(sourceDir)

    // Organize the files in the source directory.
    if err := organizer.Organize(); err != nil {
        log.Fatalf("An error occurred: %v", err)
    } else {
        fmt.Println("Folder organization completed successfully.")
    }
}
