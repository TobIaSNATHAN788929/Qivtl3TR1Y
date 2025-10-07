// 代码生成时间: 2025-10-07 19:24:52
code maintainability and scalability.
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
)

// VirtualMachine represents a virtual machine.
type VirtualMachine struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    State       string `json:"state"`
    Description string `json:"description"`
}

// virtualMachines is a slice of VirtualMachines used to store virtual machines.
var virtualMachines = []VirtualMachine{
    {
        ID:          "vm1",
        Name:        "VM-1",
        State:       "running",
        Description: "First virtual machine",
    },
    {
        ID:          "vm2",
        Name:        "VM-2",
        State:       "stopped",
        Description: "Second virtual machine",
    },
}

// GetVMs handles the GET request for virtual machines.
func GetVMs(ctx iris.Context) {
    // Send virtual machines data as JSON.
    ctx.JSON(http.StatusOK, virtualMachines)
}

// GetVM handles the GET request for a specific virtual machine by ID.
func GetVM(ctx iris.Context) {
    vmID := ctx.Params().Get("id")
    for _, vm := range virtualMachines {
        if vm.ID == vmID {
            ctx.JSON(http.StatusOK, vm)
            return
        }
    }
    ctx.JSON(http.StatusNotFound, iris.Map{"error": "Virtual machine not found"})
}

// StartVM handles the POST request to start a virtual machine.
func StartVM(ctx iris.Context) {
    vmID := ctx.Params().Get("id")
    for i, vm := range virtualMachines {
        if vm.ID == vmID {
            virtualMachines[i].State = "running"
            ctx.JSON(http.StatusOK, virtualMachines[i])
            return
        }
    }
    ctx.JSON(http.StatusNotFound, iris.Map{"error": "Virtual machine not found"})
}

// StopVM handles the POST request to stop a virtual machine.
func StopVM(ctx iris.Context) {
    vmID := ctx.Params().Get("id")
    for i, vm := range virtualMachines {
        if vm.ID == vmID {
            virtualMachines[i].State = "stopped"
            ctx.JSON(http.StatusOK, virtualMachines[i])
            return
        }
    }
    ctx.JSON(http.StatusNotFound, iris.Map{"error": "Virtual machine not found"})
}

func main() {
    app := iris.New()

    // Define routes with their respective handlers.
    app.Get("/vms", GetVMs)
    app.Get("/vms/{id}", GetVM)
    app.Post("/vms/{id}/start", StartVM)
    app.Post("/vms/{id}/stop", StopVM)

    // Start the IRIS server.
    log.Printf("Server is running at :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
