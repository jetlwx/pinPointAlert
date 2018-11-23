package models

import (
	"strings"
)

var (
	ServerURL    string
	Minutes      int
	IgnorAppName string
	IsOne        bool
	IsThree      bool
	IsFive       bool
	IsSlow       bool
	IsError      bool
	OneSum       int64
	ThreeSum     int64
	FiveSum      int64
	SlowSum      int64
	ErrorSum     int64
	LogLevel     string
	IsRecordLog  bool
	IsAlert      bool
	WXCorpID     string
	WXCorpSecret string
	WXRecver     string
	WXAgentid    int
)

//项目名
type Applications struct {
	ApplicationName string `json:"applicationName"`
	ServiceType     string `json:"serviceType"`
	Code            int    `json:"code"`
}

//单个项目的情况
type Histogram struct {
	OneS   int64 //1s
	ThreeS int64 //3s
	FiveS  int64 //5s
	Slow   int64 // gt 5s
	Error  int64 //error number
}

//最终结果
type Alert struct {
	ApplicationName string
	TotalOne        int64
	TotalThree      int64
	TotalFive       int64
	TotalSlow       int64
	TotalError      int64
}

//HasInIngnorApps is 检查单个applicationName  是否在 ignor列表中
func HasInIngnorApps(name string) bool {
	a := strings.Split(IgnorAppName, ",")
	for _, v := range a {
		if name == v {
			return true
		}
	}

	return false
}
