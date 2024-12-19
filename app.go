package main

import (
	"context"
	"math/rand"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx   context.Context
	names []string
	rng   *rand.Rand
}

// NewApp creates a new App application struct
func NewApp() *App {
	source := rand.NewSource(time.Now().UnixNano())
	return &App{
		rng: rand.New(source),
	}
}

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetNameList processes the uploaded name list and returns a slice of names
func (a *App) GetNameList(content string) []string {
	a.names = strings.Split(strings.TrimSpace(content), "\n")
	var cleanNames []string
	for _, name := range a.names {
		name = strings.TrimSpace(name)
		if name != "" {
			cleanNames = append(cleanNames, name)
		}
	}
	a.names = cleanNames
	return a.names
}

// RandomName returns a random name from the list, or a message if the list is empty
func (a *App) RandomName() string {
	if len(a.names) == 0 {
		return "请先导入名单"
	}
	return a.names[a.rng.Intn(len(a.names))]
}

