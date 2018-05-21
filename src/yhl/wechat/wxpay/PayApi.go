package wxpay

import (
	"encoding/xml"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"net/url"
	"time"
	"yhl/help"
)

func UnifiedOrder(orderReq *UnifyOrderReq) UnifyOrderResp {
	orderReq.Appid = AppId
	orderReq.Mch_id = MchId
	orderReq.Nonce_str = help.RandStr(32)

	m := help.StructToMap(*orderReq)

	orderReq.Sign = Sign(m)

	xmlByte, _ := xml.MarshalIndent(orderReq, "", "    ")
	xmlStr := string(xmlByte)

	urlStr := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	req := httplib.Post(urlStr)
	req.Header("Accept", "application/xml")
	req.Header("Content-Type", "application/xml;charset=utf-8")
	req.Body(xmlStr)
	res := UnifyOrderResp{}
	req.ToXML(&res)

	return res

}

func JsPaySdk(prepayId string) map[string]interface{} {
	m := map[string]interface{}{}

	m["appId"] = AppId
	m["timeStamp"] = time.Now().Unix()
	m["nonceStr"] = help.RandStr(32)
	m["package"] = "prepay_id=" + prepayId
	m["signType"] = "MD5"
	sign := Sign(m)
	m["paySign"] = sign

	return m
}

func GetOpenId(c *context.Context) (openid string) {
	urlStr := url.QueryEscape(c.Input.Site() + c.Input.URI())
	if code := c.Input.Query("code"); code == "" {

		codeUrl := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + AppId + "&redirect_uri=" + urlStr + "&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"

		c.Redirect(302, codeUrl)

	} else {
		openidUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + AppId + "&secret=" + AppSecret + "&code=" + code + "&grant_type=authorization_code"

		b := httplib.Get(openidUrl)
		data := make(map[string]interface{})
		b.ToJSON(&data)

		help.Log("wxpay", data)
		if v, ok := data["openid"]; ok {
			openid = v.(string)
		}
	}

	return

}
