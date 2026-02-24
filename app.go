package main

import (
	"context"
	"core-control/internal/system"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetSystemStats (internal/system/metrics.go)
func (a *App) GetSystemStats() (*system.SystemStats, error) {
	return system.GetSystemStats()
}
