package common

import "github.com/FateBug403/FateBugGUI/model/organization"

// Website 结构体
type Website struct {
	Organizations []organization.Organization `gorm:"many2many:organization_website;"` //反向关联
	Port  string `json:"port" form:"port" gorm:"column:port;comment:;size:255;"`
	Url  string `json:"url" form:"url" gorm:"column:url;comment:;size:255;"`
	Location  string `json:"location" form:"location" gorm:"column:location;comment:;size:255;"`
	Title  string `json:"title" form:"title" gorm:"column:title;comment:;size:255;collate:utf8;"`
	Scheme  string `json:"scheme" form:"scheme" gorm:"column:scheme;comment:;size:255;"`
	IP string `json:"ip" form:"ip" gorm:"column:ip;comment:;size:255;"`
	Webserver  string `json:"webserver" form:"webserver" gorm:"column:webserver;comment:;size:255;"`
	Body  string `json:"body" form:"body" gorm:"column:body;comment:;size:255;"`
	ContentType  string `json:"contentType" form:"contentType" gorm:"column:content_type;comment:;size:255;"`
	Method  string `json:"method" form:"method" gorm:"column:method;comment:;size:255;"`
	Host  string `json:"host" form:"host" gorm:"column:host;comment:;size:255;"`
	Path  string `json:"path" form:"path" gorm:"column:path;comment:;size:255;"`
	Favicon  string `json:"favicon" form:"favicon" gorm:"column:favicon;comment:;size:255;"`
	FaviconPath  string `json:"faviconPath" form:"faviconPath" gorm:"column:favicon_path;comment:;size:255;"`
	RawHeader  string `json:"rawHeader" form:"rawHeader" gorm:"column:raw_header;comment:;size:255;"`
	Time  string `json:"time" form:"time" gorm:"column:time;comment:;size:255;"`
	Jarm  string `json:"jarm" form:"jarm" gorm:"column:jarm;comment:;size:255;"`
	BodyHash  string `json:"bodyHash" form:"bodyHash" gorm:"column:body_hash;comment:;size:255;"`
	HeaderHash  string `json:"headerHash" form:"headerHash" gorm:"column:header_hash;comment:;size:255;"`
	A  string `json:"a" form:"a" gorm:"column:a;comment:;size:255;"`
	Cname  string `json:"cname" form:"cname" gorm:"column:cname;comment:;size:255;"`
	Tech  string `json:"tech" form:"tech" gorm:"column:tech;comment:;size:255;"`
	Finger string `json:"finger" form:"finger" gorm:"column:finger;comment:;size:255;"`
	Words  uint `json:"words" form:"words" gorm:"column:words;comment:;size:19;"`
	Lines  uint `json:"lines" form:"lines" gorm:"column:lines;comment:;size:19;"`
	StatusCode  uint `json:"statusCode" form:"statusCode" gorm:"column:status_code;comment:;size:19;"`
	ContentLength  uint `json:"contentLength" form:"contentLength" gorm:"column:content_length;comment:;size:19;"`
	Cdn  *bool `json:"cdn" form:"cdn" gorm:"column:cdn;comment:;"`
	ScreenshotPath  string `json:"screenshotPath" form:"screenshotPath" gorm:"column:screenshot_path;comment:;size:255;"`
	Seen *bool `json:"seen" form:"seen" gorm:"column:seen;comment:;default:false;"`
}