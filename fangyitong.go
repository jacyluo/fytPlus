package fytPlus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jacyluo/utils"
	"strings"
	"time"
)

func (e *FangYiTong) SetHeader() map[string]string {
	header := make(map[string]string)
	header["appid"] = e.Appid
	header["token"] = e.Token
	return header
}

// SendSms 发送短信
func (e *FangYiTong) SendSms(body *SmsBody) error {
	body.Sign = makeSmsSign(body, &e.Key)
	info, _ := json.Marshal(body)

	url := e.ApiUrl + "/api/v1/public/sms/send"
	res, err := utils.Client("POST", url, info, e.SetHeader())

	var obj FytRes
	if err = json.Unmarshal(res, &obj); err != nil {
		return err
	}
	if obj.Code != 200 {
		return errors.New(obj.Msg)
	}
	return nil
}

// makeSmsSign
func makeSmsSign(req *SmsBody, key *string) string {
	data := utils.BuildSignQueryStr(req.Data)
	format := "data={%s}&phone=%s&sign_name=%s&template_code=%s&timestamp=%d&key=%s"
	signStr := fmt.Sprintf(format, data, req.Phone, req.SignName, req.TemplateCode, req.Timestamp, *key)
	sign, _ := utils.GetMd5String(signStr)
	sign = strings.ToUpper(sign)
	return sign
}

//SendCode 发送验证码
func (e *FangYiTong) SendCode(req *SmsCheckBody) error {
	//smsBody := SmsCheckBody{
	//	//SignName:     "中战科技",
	//	//TemplateCode: "SMS_228845242",
	//	TemplateCode: templateCode,
	//	Phone:        phone,
	//	//CodeNum:      4,
	//	//ParamName:    "code",
	//	ExpiresIn: 600,
	//	Timestamp: time.Now().Unix(),
	//}
	url := e.ApiUrl + "/api/v1/public/sms/sendCode"

	signStr := utils.BuildSignQueryStr(utils.JSONMethod(*req))
	signStr += "&key=" + e.Key

	req.Sign, _ = utils.GetMd5String(signStr)
	req.Sign = strings.ToUpper(req.Sign)

	body, _ := json.Marshal(*req)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		var result FytRes
		if err = json.Unmarshal(res, &result); err != nil {
			return err
		}
		if result.Code != 200 {
			return errors.New(result.Msg)
		}
	}
	return nil
}

//CheckCode 验证码验证
func (e *FangYiTong) CheckCode(req *CheckCodeReq) error {
	url := e.ApiUrl + "/api/v1/public/sms/checkCode"

	type Info struct {
		Phone     string `json:"phone"`
		Code      string `json:"code"`
		Timestamp int64  `json:"timestamp"`
		Sign      string `json:"sign"`
	}
	info := Info{
		Phone:     req.Phone,
		Code:      req.Code,
		Timestamp: time.Now().Unix(),
	}

	signStr := utils.BuildSignQueryStr(utils.JSONMethod(info))
	signStr += "&key=" + e.Key

	info.Sign, _ = utils.GetMd5String(signStr)
	info.Sign = strings.ToUpper(info.Sign)

	body, _ := json.Marshal(info)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		var result FytRes
		if err = json.Unmarshal(res, &result); err != nil {
			return err
		}
		if result.Code != 200 {
			return errors.New(result.Msg)
		}
	}
	return nil
}

// GetUpToken 获取上传图片凭证
func (e *FangYiTong) GetUpToken(c *PicTokenGetReq, model *FytRes) error {
	url := e.ApiUrl + "/api/v1/public/qiniuImageToken"
	body, _ := json.Marshal(c)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		if err = json.Unmarshal(res, &model); err != nil {
			return err
		}
		if model.Code != 200 {
			return errors.New(model.Msg)
		}
	}
	return nil
}

// UpNotify 通知文件上传成功状态
func (e *FangYiTong) UpNotify(c *PicNotifyReq, model *FytRes) error {
	url := e.ApiUrl + "/api/v1/public/qiniuImageToken"
	body, _ := json.Marshal(c)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		if err = json.Unmarshal(res, &model); err != nil {
			return err
		}
		if model.Code != 200 {
			return errors.New(model.Msg)
		}
	}
	return nil
}

// Ocr 图片识别
func (e *FangYiTong) Ocr(req *OcrReq, model *FytRes) error {
	arrType := []string{"IDCardFront", "IDCardBack", "bank", "bizLicense", "generalFast"}
	if !utils.InArrayForString(arrType, req.Type) {
		return errors.New("参数types无效")
	}
	url := e.ApiUrl + "/api/v1/public/ocr"
	body, _ := json.Marshal(*req)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		if err = json.Unmarshal(res, &model); err != nil {
			return err
		}
		if model.Code != 200 {
			return errors.New(model.Msg)
		}
	}
	return nil
}

// GetRedirectUrl 获取上传图片凭证
func (e *FangYiTong) GetRedirectUrl(c *RedirectUrlReq, model *FytRes) error {
	url := e.ApiUrl + "/api/v1/public/wechatMp/getRedirectURL?state=%s&attach=%s&scope=%s"
	url = fmt.Sprintf(url, c.State, c.Attach, c.Scope)
	body, _ := json.Marshal(c)

	if res, err := utils.Client("GET", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		if err = json.Unmarshal(res, &model); err != nil {
			return err
		}
		if model.Code != 200 {
			return errors.New(model.Msg)
		}
	}
	return nil
}

// CreateOrder 创建订单
func (e *FangYiTong) CreateOrder(c *CreateOrderReq, model *FytRes) error {
	url := e.ApiUrl + "/api/v1/public/pay/orderCreate"
	body, _ := json.Marshal(*c)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		if err = json.Unmarshal(res, &model); err != nil {
			return err
		}
		if model.Code != 200 {
			return errors.New(model.Msg)
		}
	}
	return nil
}

// AiAddress 地址识别
func (e *FangYiTong) AiAddress(c *AiAddressReq, model *AiAddressRes) error {
	url := e.ApiUrl + "/api/v1/public/address"
	body, _ := json.Marshal(*c)

	if res, err := utils.Client("POST", url, body, e.SetHeader()); err != nil {
		return err
	} else {
		if err = json.Unmarshal(res, &model); err != nil {
			return err
		}
		if model.Code != 200 {
			return errors.New(model.Msg)
		}
	}
	return nil
}
