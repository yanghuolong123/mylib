package wxpay

import (
	"encoding/xml"
	"fmt"
	"strings"
	"yhl/help"
)

func UnifiedOrder(orderReq *UnifyOrderReq) {
	m := help.StructToMap(*orderReq)

	orderReq.Appid = AppId
	orderReq.Mch_id = MchId
	orderReq.Sign = Sign(m)

	orderReq.Nonce_str = help.RandStr(32)

	xmlByte, _ := xml.Marshal(orderReq)
	xmlStr := string(xmlByte)
	strXml := strings.Replace(xmlStr, "UnifyOrderReq", "xml", -1)
	fmt.Println(strXml)

}
