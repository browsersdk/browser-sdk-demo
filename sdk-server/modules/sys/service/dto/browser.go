package dto

import (
	"github.com/baowk/dilu-core/core/base"
	"github.com/browsersdk/brosdk-server-go"
)

type BrowserGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	EnvName      string `json:"envName" form:"envName"`                      //名称
	EnvId        string `json:"envId" form:"envId"`                          //环境ID
	UserId       int    `json:"userId" form:"userId" query:"column:user_id"` //用户ID
}

func (BrowserGetPageReq) TableName() string {
	return "browser"
}

// Browser
type BrowserDto struct {
	Id      int             `json:"id"`      //主键
	EnvName string          `json:"envName"` //名称
	EnvId   string          `json:"envId"`   //环境ID
	UserId  int             `json:"userId"`  //用户ID
	Data    *brosdk.EnvInfo `json:"data"`    //数据
	Status  int8            `json:"status"`  //状态 1 停止 3 启动
}

type BrowserStatusDto struct {
	Id     int    `json:"id"`     //主键
	EnvId  string `json:"envId"`  //环境ID
	UserId int    `json:"userId"` //用户ID
	Status int8   `json:"status"` //状态 1 停止 3 启动
}

type KernelVersionInfo struct {
	Id      int    `json:"id"`
	Version string `json:"version"`
}

type SYSDEVICE struct {
	Name            string `json:"name"`            //系统平台 Windows Android MAC iPhone Linux
	System          string `json:"system"`          //具体的系统版本号
	Browser         string `json:"browser"`         //浏览器上面对应的版本
	PlatformVersion string `json:"platformVersion"` //浏览器的platformVersion
}

// 操作系统和浏览器核关系
type SYSTEMKERNEL struct {
	Windows []SYSDEVICE `json:"Windows"`
	Android []SYSDEVICE `json:"Android"`
	MacOS   []SYSDEVICE `json:"MacOS"`
	IOS     []SYSDEVICE `json:"IOS"`
	Linux   []SYSDEVICE `json:"Linux"`
}

type CONAB struct {
	Country  string `json:"country"`  //国家
	Province string `json:"province"` //省
	AB       string `json:"ab"`       //简写
	ECountry string `json:"ecountry"` //国家英文名
	must     int    //是不是这个国家的必须语言 1是 0不是
	Name     string `json:"name"`
	Code     string `json:"code"`
}

type GetUiFingerList struct {
	ChromeKernelversion  []KernelVersionInfo `json:"chromeKernelVersion"`  //支持的浏览器内核大版本
	FirefoxKernelversion []KernelVersionInfo `json:"firefoxKernelversion"` //支持的浏览器火狐内核大版本
	System               SYSTEMKERNEL        `json:"system"`               //操作系统版本
	ChromeUAversion      []string            `json:"chromeUAversion"`      //浏览器UA版本
	FirefoxUAversion     []string            `json:"firefoxUAversion"`     //火狐浏览器UA版本
	Language             []CONAB             `json:"language"`             //语言
	Zone                 []string            `json:"zone"`                 //时区
	Dpi                  any                 `json:"dpi"`                  //屏幕分辨率
	Webgl                any                 `json:"webgl"`                //webgl
	Cpu                  any                 `json:"cpu"`                  //CPU参数
	Mem                  any                 `json:"mem"`                  //内存参数
	Region               any                 `json:"region"`               //Region
}
