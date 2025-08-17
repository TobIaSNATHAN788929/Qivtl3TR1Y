// 代码生成时间: 2025-08-17 20:55:38
package main

import (
    "fmt"
    "time"
    "github.com/iris-contrib/middleware/cors"
    "github.com/kataras/iris/v12"
    "github.com/robfig/cron/v3"
)

// Scheduler 定时任务调度器
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler 创建一个新的定时任务调度器
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// AddJob 添加一个新的定时任务
func (s *Scheduler) AddJob(scheduler, command string) error {
    _, err := s.cron.AddFunc(scheduler, func() { runCommand(command) })
    if err != nil {
        return err
