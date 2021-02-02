package alidayu

// MnsError 错误码
type MnsError string

func (e MnsError) Error() string {
	return string(e)
}

// DIY error
const (
	appKeyIsEmpty        = MnsError("appKey is empty")       // 空appKey
	appSecretIsEmpty     = MnsError("appSecret is empty")    // 空appSecret
	freeSignNameIsEmpty  = MnsError("freeSignName is empty") // 空freeSignName
	templateCodeIsEmpty  = MnsError("templateCode is empty") // 空templateCode
	telIsEmpty           = MnsError("tel is empty")          // 空tel
	castIsNil            = MnsError("cast is nil")           // 空cast
	getRequestBodyFailed = MnsError("getRequestBody failed") // 构造http请求失败
	castRequestFailed    = MnsError("cast request failed")   // http请求失败
)

// Alidayu error
const (
	smsSignatureSceneIllegal   = MnsError("isv.SMS_SIGNATURE_SCENE_ILLEGAL")   // 短信所使用签名场景非法
	extendCodeError            = MnsError("isv.EXTEND_CODE_ERROR")             // 扩展码使用错误，相同的扩展码不可用于多个签名
	domesticNumberNotSupported = MnsError("isv.DOMESTIC_NUMBER_NOT_SUPPORTED") // 国际/港澳台消息模板不支持发送境内号码
	denyIpRange                = MnsError("isv.DENY_IP_RANGE")                 // 源IP地址所在的地区被禁用
	dayLimitControl            = MnsError("isv.DAY_LIMIT_CONTROL")             // 触发日发送限额
	smsContentIllegal          = MnsError("isv.SMS_CONTENT_ILLEGAL")           // 短信内容包含禁止发送内容
	smsSignIllegal             = MnsError("isv.SMS_SIGN_ILLEGAL")              // 签名禁止使用
	ramPermissionDeny          = MnsError("isp.RAM_PERMISSION_DENY")           // RAM权限DENY
	outOfService               = MnsError("isv.OUT_OF_SERVICE")                // 业务停机
	productUnSubscript         = MnsError("isv.PRODUCT_UN_SUBSCRIPT")          // 未开通云通信产品的阿里云客户
	productUnsubscribe         = MnsError("isv.PRODUCT_UNSUBSCRIBE")           // 产品未开通
	accountNotExists           = MnsError("isv.ACCOUNT_NOT_EXISTS")            // 账户不存在
	accountAbnormal            = MnsError("isv.ACCOUNT_ABNORMAL")              // 账户异常
	smsTemplateIllegal         = MnsError("isv.SMS_TEMPLATE_ILLEGAL")          // 短信模版不合法
	smsSignatureIllegal        = MnsError("isv.SMS_SIGNATURE_ILLEGAL")         // 短信签名不合法
	invalidParameters          = MnsError("isv.INVALID_PARAMETERS")            // 参数异常
	systemError                = MnsError("isp.SYSTEM_ERROR")                  // 系统错误
	mobileNumberIllegal        = MnsError("isv.MOBILE_NUMBER_ILLEGAL")         // 非法手机号
	mobileCountOverLimit       = MnsError("isv.MOBILE_COUNT_OVER_LIMIT")       // 手机号码数量超过限制
	templateMissingParameters  = MnsError("isv.TEMPLATE_MISSING_PARAMETERS")   // 模版缺少变量
	businessLimitControl       = MnsError("isv.BUSINESS_LIMIT_CONTROL")        // 业务限流
	invalidJsonParam           = MnsError("isv.INVALID_JSON_PARAM")            // JSON参数不合法，只接受字符串值
	blackKeyControlLimit       = MnsError("isv.BLACK_KEY_CONTROL_LIMIT")       // 黑名单管控
	paramLengthLimit           = MnsError("isv.PARAM_LENGTH_LIMIT")            // 参数超出长度限制
	paramNotSupportUrl         = MnsError("isv.PARAM_NOT_SUPPORT_URL")         // 不支持URL
	amountNotEnough            = MnsError("isv.AMOUNT_NOT_ENOUGH")             // 账户余额不足
	templateParamsIllegal      = MnsError("isv.TEMPLATE_PARAMS_ILLEGAL")       // 模版变量里包含非法关键字
	unknownError               = MnsError("isp.UNKNOWN_ERROR")                 // 未知错误，这个错误是个大坑，官方错误码中没有这个错误码。在短信参数类型和短信模板申请的参数类型不一致时可能会出现，比如模板申请的属性是手机号，但是参数填充了英文字母，可能导致这个错误。
)

// IsSlightError 是否轻微错误（不建议报警）
func IsSlightError(err error) bool {
	switch err.Error() {
	case castRequestFailed.Error():
		return true
	case dayLimitControl.Error(), mobileNumberIllegal.Error(), businessLimitControl.Error():
		return true
	default:
		return false
	}
}

// IsRepeatableError 是否可重试错误（不建议报警，建议重试）
func IsRepeatableError(err error) bool {
	switch err.Error() {
	case castRequestFailed.Error():
		return true
	case systemError.Error():
		return true
	default:
		return false
	}
}

// IsSevereError 是否严重错误（建议报警）
func IsSevereError(err error) bool {
	return !IsSlightError(err) && !IsRepeatableError(err)
}
