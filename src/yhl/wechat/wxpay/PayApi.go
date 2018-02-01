package wxpay

import (
	"encoding/xml"
	"github.com/astaxie/beego/httplib"
	"strings"
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
	strXml := strings.Replace(xmlStr, "UnifyOrderReq", "xml", -1)

	url := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	req := httplib.Post(url)
	req.Header("Accept", "application/xml")
	req.Header("Content-Type", "application/xml;charset=utf-8")
	req.Body(strXml)
	res := UnifyOrderResp{}
	req.ToXML(&res)

	return res

}

func ToWxRespXmlStr(code, msg string) string {
	resp := new(WXPayNotifyResp)
	resp.Return_code = code
	resp.Return_msg = msg

	b, _ := xml.Marshal(resp)

	strResp := strings.Replace(string(b), "WXPayNotifyResp", "xml", -1)
	help.Log("wxpay", "strResp: "+strResp)

	return strResp
}
