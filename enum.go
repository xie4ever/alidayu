package alidayu

// 默认配置
const (
	defaultCircuitName = "alidayu_circuit"
	productionURL      = "https://eco.taobao.com/router/rest"
)

// 接口调参
const (
	messageType          = "normal"                         // 短信类型-普通
	messageMethodSms     = "alibaba.aliqin.fc.sms.num.send" // 调用接口类型-短信
	messageFormatJson    = "json"                           // 短信参数格式-json
	messageVersion       = "2.0"                            // 短信版本-2.0
	messageSignMethodMd5 = "md5"                            // 签名方式-MD5
)
