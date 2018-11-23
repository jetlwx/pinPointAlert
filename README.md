# pinPointAlert
将pinPoint采集的信息推送至微信企业号


#### Usage of ./pinPointAlert:
######   -agentid int
     	微信企业号程序ID (default 1)
  ###### -corpid string
    	微信企业号Corpid (default "xxxxxxxx")
######   -corpsecret string
    	微信企业号CorpSecret (default "xxxxxx")
 ######  -errorsum int
    	ERROR阀值,iserror 为true时生效 (default 5)
 ######  -fivesum int
    	3－5S阀值,isfive 为true时生效 (default 30)
 ######  -ignorApp string
    	不统计的项目，多个项目用逗号隔开
######   -isalert
    	是否发送微信报警
 ######  -iserror
     	是否启用ERROR统计，默认启用 (default true)
 ######  -isfive
    	是否启用大于3-5S统计，默认启用 (default true)
 ######  -isone
    	是否启用1S统计，默认不启用
 ######  -isrecoder
    	是否记录报警记录日志，默认保存运行当前目录，默认启用 (default true)
 ######  -isslow
    	是否启用大于5s统计，默认启用 (default true)
 ######  -isthree
    	是否启用1-3S统计，默认不启用
 ######  -loglevel string
    	日志级别：debug|info|warn|eror (default "eror")
 ######  -min int
    	从当前时间向前移minutes (default 5)
 ######  -onesum int
    	小于1S阀值,isone 为true时生效
######   -receiver string
    	微信接收人 (default "@all")
######   -server string
    	pinpoint服务端地址 (default "http://192.168.107.60:28080")
######   -slowsum int
    	大于5s阀值,isslow 为true时生效 (default 10)
######   -threesum int
    	1-3S阀值,isthree 为true时生效
