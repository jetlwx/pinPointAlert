package main

import (
	"flag"

	"github.com/jetlwx/comm"
	"github.com/jetlwx/pinPointAlert/models"
)

func init() {
	flag.StringVar(&models.ServerURL, "server", "http://192.168.107.60:28080", "pinpoint服务端地址")
	flag.IntVar(&models.Minutes, "min", 5, "从当前时间向前移minutes")
	flag.StringVar(&models.IgnorAppName, "ignorApp", "", "不统计的项目，多个项目用逗号隔开")
	flag.BoolVar(&models.IsOne, "isone", false, "是否启用1S统计，默认不启用")
	flag.BoolVar(&models.IsThree, "isthree", false, "是否启用1-3S统计，默认不启用")
	flag.BoolVar(&models.IsFive, "isfive", true, "是否启用大于3-5S统计，默认启用")
	flag.BoolVar(&models.IsSlow, "isslow", true, "是否启用大于5s统计，默认启用")
	flag.BoolVar(&models.IsError, "iserror", true, "是否启用ERROR统计，默认启用")
	flag.Int64Var(&models.OneSum, "onesum", 0, "小于1S阀值,isone 为true时生效")
	flag.Int64Var(&models.ThreeSum, "threesum", 0, "1-3S阀值,isthree 为true时生效")
	flag.Int64Var(&models.FiveSum, "fivesum", 30, "3－5S阀值,isfive 为true时生效")
	flag.Int64Var(&models.SlowSum, "slowsum", 10, "大于5s阀值,isslow 为true时生效")
	flag.Int64Var(&models.ErrorSum, "errorsum", 5, "ERROR阀值,iserror 为true时生效")
	flag.StringVar(&models.LogLevel, "loglevel", "eror", "日志级别：debug|info|warn|eror")
	flag.BoolVar(&models.IsRecordLog, "isrecoder", true, "是否记录报警记录日志，默认保存运行当前目录，默认启用")
	flag.BoolVar(&models.IsAlert, "isalert", false, "是否发送微信报警")
	flag.StringVar(&models.WXCorpID, "corpid", "xxxxxxxx", "微信企业号Corpid")
	flag.StringVar(&models.WXCorpSecret, "corpsecret", "xxxxxx", "微信企业号CorpSecret")
	flag.StringVar(&models.WXRecver, "receiver", "@all", "微信接收人")
	flag.IntVar(&models.WXAgentid, "agentid", 1, "微信企业号程序ID")
	flag.Parse()
	comm.LogLevel = models.LogLevel
}
func main() {

	models.Do()
}
