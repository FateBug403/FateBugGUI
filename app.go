package main

import (
	"github.com/FateBug403/FateBugGUI/api"
	afrogReq "github.com/FateBug403/FateBugGUI/model/afrog/request"
	"github.com/FateBug403/FateBugGUI/model/common"
	"context"
	"fmt"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.Ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// InfoCollectionStart test
func (a *App) InfoCollectionStart()string{
	api.InfoCollectionApi.Start(a.Ctx)
	return "test"
}
// GetAfrogList 获取Afrog POC列表
func (a *App) GetAfrogList(SearchInfo afrogReq.PocSearchInfo) common.H {
	return api.AfrogAPi.GetAfrogList(SearchInfo)
}

// ReadAfrogPoc // 模糊搜索Afrog POC
func (a *App) ReadAfrogPoc(pocId string) common.H{
	return api.AfrogAPi.ReadAfrogPoc(pocId)
}