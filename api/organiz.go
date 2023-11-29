package api

import (
	"github.com/FateBug403/FateBugGUI/model/common"
	"github.com/FateBug403/FateBugGUI/model/organization"
	organizReq "github.com/FateBug403/FateBugGUI/model/organization/request"
)

type organizAPi struct {

}

// GetOrganization 获取组织列表
func (receiver organizAPi)GetOrganization(SearchInfo organizReq.OrganizationSearch) common.H {
	return common.M{"code":200,"date":""}
}
// CreateOrganization 添加组织
func (receiver organizAPi)CreateOrganization(parameter organization.Organization) common.H {
	return common.M{"code":200,"date":""}
}

// DeleteOrganizationByIds 通过Id删除特定组织
func (receiver organizAPi) DeleteOrganizationByIds(parameter common.IdsReq) common.H{
	return common.M{"code":200,"date":""}
}

var OrganizAPi = new(organizAPi)