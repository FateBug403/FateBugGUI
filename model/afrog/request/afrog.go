package request

import (
	"github.com/FateBug403/FateBugGUI/model/afrog"
	"time"
)

type AfrogSearch struct{
    afrog.Afrog
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
}

type PocSearchInfo struct {
	Keyword string //通过关键词搜索(name和id中模糊匹配)
	Description string // 通过描述搜索关键字
}