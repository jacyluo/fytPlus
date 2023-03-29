package main

import (
	"github.com/jacyluo/fytPlus"
	"time"
)

func main() {

}

type SendSmsReq struct {
	Phone string `json:"phone"`
}

func SendCode(req *SendSmsReq) error {
	fyt := fytPlus.FangYiTong{
		ApiUrl: "ApiUrl",
		Appid:  "Appid",
		Token:  "Token",
		Key:    "Key",
	}
	body := fytPlus.SmsCheckBody{
		Phone:        req.Phone,
		SignName:     "中战科技",
		TemplateCode: "SMS_228845242",
		//CodeNum:      4, // 默认值 4
		//ParamName:    "code", // 默认值 code
		//ExpiresIn: 600, // 默认值 600
		Timestamp: time.Now().Unix(), // 如果时间差超过60分钟，则报错
	}

	if err := fyt.SendCode(&body); err != nil {
		return err
	}
	return nil
}

func CheckCode(req *fytPlus.CheckCodeReq) error {
	fyt := fytPlus.FangYiTong{
		ApiUrl: "ApiUrl",
		Appid:  "Appid",
		Token:  "Token",
		Key:    "Key",
	}
	if err := fyt.CheckCode(req); err != nil {
		return err
	}
	return nil
}

func SendSms(req *SendSmsReq) error {
	fyt := fytPlus.FangYiTong{
		ApiUrl: "ApiUrl",
		Appid:  "Appid",
		Token:  "Token",
		Key:    "Key",
	}
	param := make(map[string]interface{})

	templateCode := "SMS_228845242" // 验证码
	param["code"] = "6789"

	//templateCode := "SMS_169101558" //没有参数

	//templateCode := "SMS_228845348" //两个参数
	//param["prodName"] = "儿童手表"
	//param["cancelTime"] = time.Now().Format("2006-01-02 15:04:05")

	body := fytPlus.SmsBody{
		Phone:        req.Phone,
		SignName:     "中战科技",
		TemplateCode: templateCode,
		Data:         param,
		Timestamp:    time.Now().Unix(), // 如果时间差超过60分钟，则报错
	}
	if err := fyt.SendSms(&body); err != nil {
		return err
	}
	return nil
}

func Ocr(req *fytPlus.OcrReq, model *fytPlus.FytRes) error {
	fyt := fytPlus.FangYiTong{
		ApiUrl: "ApiUrl",
		Appid:  "Appid",
		Token:  "Token",
		Key:    "Key",
	}
	if err := fyt.Ocr(req, model); err != nil {
		return err
	}
	return nil
}

// UpTokenGet 获取图片上传token
func UpTokenGet(c *fytPlus.RedirectUrlReq, model *fytPlus.FytRes) error {
	var s fytPlus.FangYiTong
	s.GetRedirectUrl(c, model)
	return nil
}
