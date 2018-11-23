package models

import (
	"encoding/json"
	"fmt"

	"strconv"

	"time"

	"github.com/jetlwx/comm"
	"github.com/jetlwx/comm/fileAction"
	"github.com/jetlwx/comm/httplib"
	"github.com/jetlwx/comm/weixinAPI"
)

func Do() {
	var alertContent string
	t := time.Now().Format("2006-01-02 15:04:05") + " "
	//1,get all applicationName list
	apps := GetApplications()
	for _, app := range apps {
		var str string
		var has bool
		a := AppRes(app)
		var ac string
		//2. it is has in ignor list
		if HasInIngnorApps(a.ApplicationName) {
			comm.JetLog("I", a.ApplicationName, "被忽略")
			continue
		}

		if IsOne && a.TotalOne >= OneSum {
			res := "1s之内：" + strconv.FormatInt(a.TotalOne, 10) + "个 "
			str += res
			has = true

			ac += res + "\n"

		}
		if IsThree && a.TotalThree >= ThreeSum {
			res := "1-3s：" + strconv.FormatInt(a.TotalThree, 10) + "个 "
			str += res
			has = true
			ac += res + "\n"
		}
		if IsFive && a.TotalFive >= FiveSum {
			res := "3-5s：" + strconv.FormatInt(a.TotalFive, 10) + "个 "
			str += res
			has = true
			ac += res + "\n"

		}
		if IsSlow && a.TotalSlow >= SlowSum {
			res := "大于5s：" + strconv.FormatInt(a.TotalSlow, 10) + "个 "
			str += res
			has = true
			ac += res + "\n"
		}
		if IsError && a.TotalError >= ErrorSum {
			res := "访问出错：" + strconv.FormatInt(a.TotalError, 10) + "个 "
			str += res
			has = true
			ac += res + "\n"
		}

		if has {

			str1 := a.ApplicationName + " "
			str1 += str
			comm.JetLog("I", "模块计算情况：", str1, "\n")

			if IsRecordLog {
				RecordLog(t + "模块=" + a.ApplicationName + "  1s之内=" + strconv.FormatInt(a.TotalOne, 10) + "  1-3s=" + strconv.FormatInt(a.TotalThree, 10) + "  3-5s=" + strconv.FormatInt(a.TotalFive, 10) + "  5s以上=" + strconv.FormatInt(a.TotalSlow, 10) + "  错误数=" + strconv.FormatInt(a.TotalError, 10) + "\n")
			}

			if IsAlert && ac != "" {
				c := "模块:" + a.ApplicationName + "\n"
				c += ac + "---------------------" + "\n"
				alertContent += c
			}
		}

	}

	fmt.Println("alertContent=", alertContent)
	if IsAlert && alertContent != "" {
		c := "pinPoint访问时长报警：" + "\n"
		c += "时间：" + t + "\n" + "---------------------" + "\n"
		c += alertContent
		SendWX(c)
	}
}

//GetApplications 获取所有项目名
func GetApplications() (app []Applications) {
	url := ServerURL + "/applications.pinpoint"
	res := httplib.Get(url)
	resp, err := res.Response()
	if err != nil {
		comm.JetLog("E", err)
		return
	}

	if resp.StatusCode != 200 {
		comm.JetLog("E", "返回值;", resp.StatusCode)
	}

	body, err := res.Bytes()
	if err != nil {
		comm.JetLog("E", err)
		return
	}

	json.Unmarshal(body, &app)

	return
}

//RecordLog is 记录日志
func RecordLog(content string) error {
	f := fileAction.FileORDirPath{}
	f.Path = "alert_log.log"

	return f.AppendToFileEnd(content)
}

func SendWX(wxmsg string) {
	wx := weixinAPI.CorpInfo{}
	wx.CorpID = WXCorpID
	wx.CorpSecret = WXCorpSecret
	msg := weixinAPI.SendMsg{}
	msg.Touser = WXRecver
	msg.Agentid = WXAgentid
	msg.Text.Content = wxmsg

	wx.Send(msg)
}
