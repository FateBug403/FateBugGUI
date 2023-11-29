package main

import (
	"context"
	"fmt"
	"github.com/FateBug403/FateBugGUI/api"
	afrogReq "github.com/FateBug403/FateBugGUI/model/afrog/request"
	"github.com/FateBug403/FateBugGUI/model/common"
	"github.com/FateBug403/FateBugGUI/model/organization"
	organizReq"github.com/FateBug403/FateBugGUI/model/organization/request"
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

// GetOrganization 获取Web组织列表
func (a *App) GetOrganization(SearchInfo organizReq.OrganizationSearch) common.H {
	return api.OrganizAPi.GetOrganization(SearchInfo)
}
// CreateOrganization 获取Web组织列表
func (a *App) CreateOrganization(parameter organization.Organization) common.H {
	return api.OrganizAPi.CreateOrganization(parameter)
}
// DeleteOrganizationByIds 获取Web组织列表
func (a *App) DeleteOrganizationByIds(parameter common.IdsReq) common.H {
	return api.OrganizAPi.DeleteOrganizationByIds(parameter)
}