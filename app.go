package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	db  *DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.db = NewDB()
}

func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetAllNades(mapId string) []Nade {
	r := a.db.FindNades(mapId)

	return r
}

func (a *App) GetAllVariants(nadeId string) []Variant {
	return a.db.FindVariants(nadeId)
}

func (a *App) InsertNade(nade Nade) string {
	return a.db.InsertNade(nade)
}

func (a *App) InsertVariant(variant Variant) string {
	return a.db.InsertVariant(variant)
}

func (a *App) UpdateNade(nade Nade) {
	a.db.UpdateNade(nade)
}

func (a *App) UpdateVariant(variant Variant) {
	a.db.UpdateVariant(variant)
}

func (a *App) DeleteNade(nadeId string) {
	a.db.DeleteNade(nadeId)
}

func (a *App) DeleteVariant(variantId string) {
	a.db.DeleteVariant(variantId)
}
