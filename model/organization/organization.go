package organization

import "github.com/FateBug403/FateBugGUI/model/common"

// Organization 结构体
type Organization struct {
	ID   uint `json:"id" form:"id" gorm:"column:id;comment:;size:19;"`
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`
	Description  string `json:"description" form:"description" gorm:"column:description;comment:;size:255;"`
	Websites []common.Website `gorm:"many2many:organization_website;"` //反向关联
}
