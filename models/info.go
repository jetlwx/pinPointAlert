package models

import (
	"time"

	"strconv"

	"encoding/json"

	"github.com/bitly/go-simplejson"
	"github.com/jetlwx/comm"
	"github.com/jetlwx/comm/httplib"
)

//AppRes is 获取单个项目运行情况
func AppRes(app Applications) (a Alert) {
	body, err := info(app)
	if err != nil {
		comm.JetLog("E", err)
		return
	}

	js, err := simplejson.NewJson(body)
	if err != nil {
		comm.JetLog("E", err)
		return
	}

	linkData, err := js.Get("applicationMapData").Get("linkDataArray").Array()
	if err != nil {
		comm.JetLog("E", err)
		return
	}

	var tOne, tThree, tFive, tSlow, tError int64
	for _, v := range linkData {
		v1, _ := json.Marshal(v)
		js2, err := simplejson.NewJson(v1)
		if err != nil {
			comm.JetLog("E", err)
			continue
		}
		r, err := js2.Get("histogram").Map()
		if err != nil {
			continue
		}

		if one, ok := r["1s"].(json.Number); ok {
			if o, err := one.Int64(); err == nil {
				tOne += o
			}
		}

		if three, ok := r["3s"].(json.Number); ok {
			if t, err := three.Int64(); err == nil {
				tThree += t
			}

		}
		if five, ok := r["5s"].(json.Number); ok {
			if f, err := five.Int64(); err == nil {
				tFive += f
			}

		}

		if slow, ok := r["Slow"].(json.Number); ok {
			if s, err := slow.Int64(); err == nil {
				tSlow += s
			}

		}
		if errn, ok := r["Error"].(json.Number); ok {
			if e, err := errn.Int64(); err == nil {
				tError += e
			}

		}

	}

	a.ApplicationName = app.ApplicationName
	a.TotalError = tError
	a.TotalFive = tFive
	a.TotalOne = tOne
	a.TotalSlow = tSlow
	a.TotalThree = tThree

	return a
}

////http://192.168.107.60:28080/getServerMapData.pinpoint?applicationName=coa-im-message&from=1541757493000&to=1541757793000&callerRange=1&calleeRange=1&bidirectional=false&wasOnly=false&serviceTypeName=TOMCAT&_=1541758151830
//info is get an applicaton infomation body
func info(app Applications) (body []byte, err error) {
	tnow := time.Now().Unix() * 1000
	t5mago := tnow - int64(Minutes*60*1000)
	comm.JetLog("D", "Now=", tnow, Minutes, "ago=", t5mago)

	url := ServerURL + "/getServerMapData.pinpoint?applicationName=" + app.ApplicationName + "&"
	url += "from=" + strconv.FormatInt(t5mago, 10) + "&"
	url += "to=" + strconv.FormatInt(tnow, 10) + "&"
	url += "callerRange=1&calleeRange=1&bidirectional=false&wasOnly=false&"
	url += "serviceTypeName=" + app.ServiceType
	comm.JetLog("D", "url=", url)

	res := httplib.Get(url)
	resp, err := res.Response()
	if err != nil {
		comm.JetLog("E", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		comm.JetLog("E", "返回值;", resp.StatusCode)
	}

	body, err = res.Bytes()
	if err != nil {
		comm.JetLog("E", err)
		return nil, err
	}

	return
}
