package fytPlus

type FangYiTong struct {
	ApiUrl string `json:"apiUrl"`
	Appid  string `json:"appid"`
	Token  string `json:"token"`
	Key    string `json:"key"`
}

// SmsBody 通用短信模版
type SmsBody struct {
	SignName     string                 `json:"sign_name"`
	TemplateCode string                 `json:"template_code"`
	Phone        string                 `json:"phone" vd:"phone($,'CN'); msg:sprintf('手机号无效');"`
	Timestamp    int64                  `json:"timestamp"`
	Sign         string                 `json:"sign"`
	Data         map[string]interface{} `json:"data"` // 参数列表
}

//CheckCodeReq 验证码验证
type CheckCodeReq struct {
	Phone string `json:"phone" vd:"phone($,'CN'); msg:sprintf('手机号无效');"`
	Code  string `json:"code" vd:"regexp('^[0-9a-zA-Z]{4,6}$'); msg:sprintf('验证码无效');"`
}

type SmsCheckBody struct {
	SignName     string `json:"sign_name"'`
	TemplateCode string `json:"template_code"`
	Phone        string `json:"phone" vd:"phone($,'CN'); msg:sprintf('手机号无效');"`
	ExpiresIn    int64  `json:"expires_in"`
	ParamName    string `json:"param_name"`
	CodeNum      int    `json:"code_num"`
	Timestamp    int64  `json:"timestamp"`
	Sign         string `json:"sign"`
}

// FytRes 方蚁通通用返回值
type FytRes struct {
	RequestId string      `json:"requestId"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

type PicTokenGetReq struct {
	Domain string   `json:"domain" vd:"len($)==0 || chkUrl($); msg:sprintf('invalid domain')"`
	Bucket string   `json:"bucket" vd:"len($)==0 || regexp('^\\w[-\\w]{2,}$'); msg:sprintf('invalid bucket')"`
	List   []string `json:"list" vd:"range($,regexp('^.*?\\.(?i)(jpg|gif|png|jpeg)$',#v)); msg:sprintf('图片文件名无效')"`
	//Flag string   `json:"flag"`
}

type PicNotifyReq struct {
	List []string `json:"list" vd:"range($,regexp('^\\{23,26}$',#v)); msg:sprintf('upload_id 无效')"`
}

//OcrReq
//IDCardFront	身份证正面照
//IDCardBack	身份证背面照
//bank	银行卡
//bizLicense	营业执照
//generalFast	通用印刷文字
type OcrReq struct {
	Type string `json:"type" vd:"chkEn($,4,30);msg:sprintf('参数type无效')"`
	Url  string `json:"url" vd:"regexp('^http(s)?://.{8,}$');msg:sprintf('参数url无效')"`
}

type RedirectUrlReq struct {
	State  string `json:"state" comment:"会直接返回给客户"`
	Attach string `json:"attach" comment:"客户回调URL别名"`
	Scope  string `json:"scope" comment:"snsapi_base | snsapi_userinfo"`
}

type CreateOrderReq struct {
	ChannelId   int    `json:"channel_id"`
	ChannelType int    `json:"channel_type"`
	PayType     int    `json:"pay_type"`
	IsFrozen    int    `json:"is_frozen"`
	TradeType   string `json:"trade_type"`
	MchId       string `json:"mch_id"`
	TotalFee    string `json:"total_fee"`
	Body        string `json:"body"`
	TrxNo       string `json:"trx_no"`
	NonceStr    string `json:"nonce_str"`
	CreateIp    string `json:"create_ip"`
	NotifyUrl   string `json:"notify_url"`
	SignType    string `json:"sign_type"`
	Openid      string `json:"openid"`
	OpenidToken string `json:"openid_token"`
	Attach      string `json:"attach"`
	Sign        string `json:"sign"`
}

type AiAddressReq struct {
	Text string `json:"text" vd:"len($)>=10;msg:sprintf('要自动识别的地址不得少于10个字符')"`
}

type AiAddressRes struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      struct {
		Lat          float64 `json:"lat"`
		Detail       string  `json:"detail"`
		Town         string  `json:"town"`
		Phonenum     string  `json:"phonenum"`
		CityCode     string  `json:"city_code"`
		Province     string  `json:"province"`
		Person       string  `json:"person"`
		Lng          float64 `json:"lng"`
		ProvinceCode string  `json:"province_code"`
		Text         string  `json:"text"`
		County       string  `json:"county"`
		City         string  `json:"city"`
		CountyCode   string  `json:"county_code"`
		TownCode     string  `json:"town_code"`
		LogId        int64   `json:"log_id"`
	} `json:"data"`
}
