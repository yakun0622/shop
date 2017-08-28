package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

var smsSID, smsAppID, smsAuthToken string

func init() {
	smsSID = beego.AppConfig.DefaultString("SmsSid", "")
	smsAppID = beego.AppConfig.DefaultString("SmsAppId", "")
	smsAuthToken = beego.AppConfig.DefaultString("SmsAuthToken", "")
}

func SendSms(phone string, templateID string, param string) (string, error) {
	//格式化时间
	tm := time.Now().Format("20060102030405")
	//生成SigParameter，MD5
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(smsSID + smsAuthToken + tm))
	signParameter := strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
	//生成Authorization，Base64
	authorization := base64.StdEncoding.EncodeToString([]byte(smsSID + ":" + tm))

	postJSON := `{"templateSMS":{"appId":"#appId#","param":"#param#","templateId":"#templateId#","to":"#to#"}}`
	postJSON = strings.Replace(postJSON, "#appId#", smsAppID, 1)
	postJSON = strings.Replace(postJSON, "#param#", param, 1)
	postJSON = strings.Replace(postJSON, "#templateId#", templateID, 1)
	postJSON = strings.Replace(postJSON, "#to#", phone, 1)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", "https://api.ucpaas.com/2014-06-30/Accounts/"+
		smsSID+"/Messages/templateSMS?sig="+signParameter, strings.NewReader(postJSON))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Add("Authorization", authorization)

	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		return bodystr, nil
	}
	return "", errors.New("网络请求错误，代码：" + strconv.Itoa(response.StatusCode))
}
