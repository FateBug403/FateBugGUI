package api

import (
	command "github.com/FateBug403/FateBugGUI/commonfun"
	afrogReq "github.com/FateBug403/FateBugGUI/model/afrog/request"
	afrogRsp "github.com/FateBug403/FateBugGUI/model/afrog/response"
	"github.com/FateBug403/FateBugGUI/model/common"
	"github.com/FateBug403/FateBugGUI/pkg/afrog/pocs"
	"github.com/projectdiscovery/gologger"
	"github.com/zan8in/afrog/v2/pkg/poc"
)

type afrogAPi struct {

}

func (receiver afrogAPi)GetAfrogList(SearchInfo afrogReq.PocSearchInfo) common.H {

	// 获取全部poc
	var PocList []afrogRsp.PocListResponse
	var number = 1

	if len(pocs.EmbedFileList) > 0 {
		gologger.Print().Msg("---------- Embed PoCs -----------------")
		for _, v := range pocs.EmbedFileList {
			if pocTmp, err := pocs.EmbedReadPocByPath(v); err == nil {
				PocList = append(PocList,afrogRsp.PocListResponse{
					Index:       number,
					Id:          pocTmp.Id,
					Name:        pocTmp.Info.Name,
					Severity:    pocTmp.Info.Severity,
					Description: pocTmp.Info.Description,
				})
				number++
			}
		}
	}

	// init LocalPocsDirectory
	if len(poc.LocalFileList) > 0 {
		gologger.Print().Msg("---------- Local afrog-pocs -----------------")
		for _, v := range poc.LocalFileList {
			if pocTmp, err := poc.LocalReadPocByPath(v); err == nil {
				PocList = append(PocList,afrogRsp.PocListResponse{
					Index:       number,
					Id:          pocTmp.Id,
					Name:        pocTmp.Info.Name,
					Severity:    pocTmp.Info.Severity,
					Description: pocTmp.Info.Description,
				})
				number++
			}
		}
	}

	// append pocs
	if len(poc.LocalAppendList) > 0 {
		gologger.Print().Msg("---------- Local append-pocs -----------------")
		for _, v := range poc.LocalAppendList {
			if pocTmp, err := poc.LocalReadPocByPath(v); err == nil {
				PocList = append(PocList,afrogRsp.PocListResponse{
					Index:       number,
					Id:          pocTmp.Id,
					Name:        pocTmp.Info.Name,
					Severity:    pocTmp.Info.Severity,
					Description: pocTmp.Info.Description,
				})
				number++
			}
		}
	}

	var PocListResult []afrogRsp.PocListResponse
	for _,Poc := range PocList{
		if ((SearchInfo.Keyword == "" || command.InfoContains(Poc.Id, SearchInfo.Keyword)) || command.InfoContains(Poc.Name,SearchInfo.Keyword)) &&
			(SearchInfo.Description == "" || command.InfoContains(Poc.Description, SearchInfo.Description)){
			PocListResult = append(PocListResult, Poc)
		}
	}


	return common.M{"code":200,"date":PocListResult}
}

func (receiver afrogAPi) ReadAfrogPoc(pocId string) common.H {
	date,err:=pocs.EmbedReadContentByName(pocId)
	if err != nil {
		return common.M{"code":500,"date":"错误的读取文件"}
	}
	//log.Println(string(date))
	return common.M{"code":200,"date":string(date)}
}

var AfrogAPi = new(afrogAPi)