package wxpay

import (
	"fmt"
	"sort"
	"strings"
	"yhl/help"
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
	Product_id       int    `xml:"product_id"`       // 商品id
	Attach           string `xml:"attach"`           //
	Goods_tag        string `xml:"goods_tag"`        //
	Time_start       string `xml:"time_start"`
	Time_expire      string `xml:"time_expire"`
}

type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
	Code_url    string `xml:"code_url"`
}

func Sign(m map[string]interface{}) string {
	// 对key进行排序
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// 对key=value的键值对用&连接起来，略过空值
	var signStr string
	for _, k := range keys {
		v := fmt.Sprintf("%v", m[k])
		if v != "" {
			signStr += k + "=" + v + "&"
		}
	}

	// 在string后加入KEY
	signStr += "key=" + Key

	// MD5加密并转化为大写
	cipherStr := help.Md5(signStr)
	upperSign := strings.ToUpper(cipherStr)

	return upperSign
}

type WXPayNotifyReq struct {
	Return_code string `xml:"return_code"`
	//	Return_msg     string `xml:"return_msg"`
	Appid        string `xml:"appid"`
	Mch_id       string `xml:"mch_id"`
	Nonce_str    string `xml:"nonce_str"`
	Sign         string `xml:"sign"`
	Result_code  string `xml:"result_code"`
	Openid       string `xml:"openid"`
	Is_subscribe string `xml:"is_subscribe"`
	Trade_type   string `xml:"trade_type"`
	Bank_type    string `xml:"bank_type"`
	Total_fee    int    `xml:"total_fee"`
	Fee_type     string `xml:"fee_type"`
	Cash_fee     int    `xml:"cash_fee"`
	//	Cash_fee_Type  string `xml:"cash_fee_type"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	//	Attach         string `xml:"attach"`
	Time_end string `xml:"time_end"`
}

type WXPayNotifyResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}
