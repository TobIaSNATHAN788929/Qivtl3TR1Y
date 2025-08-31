// 代码生成时间: 2025-08-31 15:47:33
package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/irislogger"
    "github.com/kataras/iris/v12/middleware/irisrecover"
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// DatabaseConfig holds the configuration for the database
type DatabaseConfig struct {
    DSN string
}

func main() {
    // Set up the Iris application
    app := iris.New()
    app.Use(irisrecover.New())
    app.Use(irislogger.New())

    // Configure the database
    dbConfig := DatabaseConfig{DSN: "file:./migration.db?mode=rwc&cache=shared&_foreign_keys=true"}
    db, err := gorm.Open(sqlite.Open(dbConfig.DSN), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // Migrate the database
    err = migrateDatabase(db)
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    // Define routes
    app.Get("/migrate", func(ctx iris.Context) {
        ctx.JSON(iris.StatusOK, iris.Map{
            "message": "Database migration successful",
        })
    })

    // Start the server
    app.Listen(":8080")
}

// Migrate the database using gormigrate
func migrateDatabase(db *gorm.DB) error {
    // Define migrations
    migrations := []*gormigrate.Migration{
        {
            ID: "init", // Unique ID for the migration
            Migrate: func(tx *gorm.DB) error {
                // Create table
                err := tx.AutoMigrate(&User{}) // Replace with your model
                if err != nil {
                    return err
                }
                return nil
            },
             Rollback: func(tx *gorm.DB) error {
                // Drop table
                err := tx.Migrator().DropTable(&User{}) // Replace with your model
                if err != nil {
                    return err
                }
                return nil
            },
        },
    }

    // Initialize the gormigrate manager
    gormigrateManager := gormigrate.New(tx, migrations...)

    // Perform the migration
    if err := gormigrateManager.Migrate(); err != nil {
        return fmt.Errorf("failed to migrate: %w", err)
    }

    return nil
}

// User represents a user in the database
// This should be replaced with your actual model
type User struct {
    gorm.Model
    Name string
    Email string
}
