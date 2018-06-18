package wxpay

import (
	"crypto/tls"
	"encoding/xml"
	"github.com/astaxie/beego/httplib"
	"yhl/help"
)

//付款订单
type WithdrawOrder struct {
	XMLName   xml.Name `xml:"xml"`
	Mch_appid string   `xml:"mch_appid"`
	Mchid     string   `xml:"mchid"`
	//DeviceInfo     string   `xml:"device_info"`
	Nonce_str        string `xml:"nonce_str"`
	Partner_trade_no string `xml:"partner_trade_no"`
	Openid           string `xml:"openid"`
	Check_name       string `xml:"check_name"`
	Amount           int    `xml:"amount"`
	Desc             string `xml:"desc"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Sign             string `xml:"sign"`
}

//付款订单结果
type WithdrawResult struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code"`
	ErrCodeDes     string `xml:"err_code_des"`
	PaymentNo      string `xml:"payment_no"`
	PartnerTradeNo string `xml:"partner_trade_no"`
}

func PayToUser(amount float64, openid, partnerTradeNo, desc, clientIp, crtPath string) WithdrawResult {
	order := WithdrawOrder{}
	order.Mch_appid = AppId
	order.Mchid = MchId
	order.Openid = openid
	order.Amount = int(amount * 100)
	order.Desc = desc
	order.Partner_trade_no = partnerTradeNo
	//order.DeviceInfo = "WEB"
	order.Check_name = "NO_CHECK" //NO_CHECK：不校验真实姓名 FORCE_CHECK：强校验真实姓名
	order.Spbill_create_ip = clientIp
	order.Nonce_str = help.RandStr(32)
	m := help.StructToMap(order)
	order.Sign = Sign(m)
	help.Log("wxpay", order)
	xmlByte, _ := xml.MarshalIndent(order, "", "    ")
	xmlStr := string(xmlByte)
	help.Log("wxpay", "xmlStr:"+xmlStr)

	urlStr := "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	req := httplib.Post(urlStr)
	req.Header("Accept", "application/xml")
	req.Header("Content-Type", "application/xml;charset=utf-8")
	//req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	cer, err := tls.LoadX509KeyPair(crtPath+"/apiclient_cert.pem", crtPath+"/apiclient_key.pem")
	//cer, err := tls.LoadX509KeyPair(crtPath + "/apiclient_cert.p12")
	help.Error(err)
	req.SetTLSClientConfig(&tls.Config{Certificates: []tls.Certificate{cer}})
	req.Body(xmlStr)
	var res WithdrawResult
	res = WithdrawResult{}
	req.ToXML(&res)

	return res
}
