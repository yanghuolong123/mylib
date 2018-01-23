package wxpay

import (
	"fmt"
)

type UnifyOrderReq struct {
	Appid            string `xml:"appid"`            //公众账号ID
	Body             string `xml:"body"`             //商品描述
	Mch_id           string `xml:"mch_id"`           //商户号
	Nonce_str        string `xml:"nonce_str"`        //随机字符串
	Notify_url       string `xml:"notify_url"`       //通知地址
	Trade_type       string `xml:"trade_type"`       //交易类型
	Spbill_create_ip string `xml:"spbill_create_ip"` //支付提交用户端ip
	Total_fee        int    `xml:"total_fee"`        //总金额
	Out_trade_no     string `xml:"out_trade_no"`     //商户订单号
	Sign             string `xml:"sign"`             //签名
	Openid           string `xml:"openid"`           //购买商品的用户wxid
}
