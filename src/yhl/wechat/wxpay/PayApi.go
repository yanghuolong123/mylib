package wxpay

import (
	"encoding/xml"
	"github.com/astaxie/beego/httplib"
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

	url := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	req := httplib.Post(url)
	req.Header("Accept", "application/xml")
	req.Header("Content-Type", "application/xml;charset=utf-8")
	req.Body(xmlStr)
	res := UnifyOrderResp{}
	req.ToXML(&res)

	return res

}
