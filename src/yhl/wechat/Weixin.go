package wechat

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"io"
	"sort"
	"strings"
	"time"
	"yhl/help"
	"yhl/wechat/wxpay"
)

const (
	ApiUrl = "https://api.weixin.qq.com"
)

var (
	Token  string
	Appid  string
	Secret string
)

func init() {
	Token = beego.AppConfig.String("wechat.token")
	Appid = beego.AppConfig.String("wechat.appid")
	Secret = beego.AppConfig.String("wechat.secret")
}

type MsgBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
	Event        string
	EventKey     string
}

func Check(timestamp, nonce, signatureIn string) bool {
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	signatureGen := fmt.Sprintf("%x", s.Sum(nil))

	return signatureGen == signatureIn
}

func GetAccessToken() (token string) {
	cache := help.Cache
	t := cache.Get("access_token_" + Appid)
	if t != nil {
		token = string(t.([]uint8))
		return
	}

	url := ApiUrl + "/cgi-bin/token?grant_type=client_credential&appid=" + Appid + "&secret=" + Secret

	b := httplib.Get(url)
	data := make(map[string]interface{})
	b.ToJSON(&data)
	fmt.Println(data)

	if v, ok := data["access_token"]; ok {
		token = v.(string)
		ttl := time.Duration(data["expires_in"].(float64))
		cache.Put("access_token_"+Appid, token, ttl*time.Second)
	}

	return
}

func SendMsg(m map[string]interface{}) {
	url := ApiUrl + "/cgi-bin/message/custom/send?access_token=" + GetAccessToken()
	req := httplib.Post(url)
	//fmt.Println(m)
	req.JSONBody(m)
	req.String()
	//fmt.Println(req.String())
}

func SendTextMsg(touser, content string) {
	m := map[string]interface{}{}
	m["touser"] = touser
	m["msgtype"] = "text"
	m["text"] = map[string]string{"content": content}

	SendMsg(m)
}

func GetWxUserinfo(openid, lang string) (m map[string]interface{}) {
	if lang == "" {
		lang = "zh_CN"
	}
	url := ApiUrl + "/cgi-bin/user/info?access_token=" + GetAccessToken() + "&openid=" + openid + "&lang=" + lang
	req := httplib.Get(url)
	m = make(map[string]interface{})
	req.ToJSON(&m)

	return
}

func GetQrCodeImg(m map[string]interface{}) (imgUrl string) {
	url := ApiUrl + "/cgi-bin/qrcode/create?access_token=" + GetAccessToken()
	req := httplib.Post(url)
	req.JSONBody(m)
	data := make(map[string]interface{})
	req.ToJSON(&data)

	if v, ok := data["ticket"]; ok {
		imgUrl = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + v.(string)
	}

	return
}

func GetTmpStrQrImg(sceneStr string) (imgUrl string) {
	cache := help.Cache
	t := cache.Get("qr_img_" + sceneStr)
	if t != nil {
		imgUrl = string(t.([]uint8))
		return
	}

	expire := 1800
	m := map[string]interface{}{}
	m["action_name"] = "QR_STR_SCENE"
	m["action_info"] = map[string]interface{}{"scene": map[string]string{"scene_str": sceneStr}}
	m["expire_seconds"] = expire

	imgUrl = GetQrCodeImg(m)
	if len(imgUrl) > 0 {
		cache.Put("qr_img_"+sceneStr, imgUrl, time.Duration(expire)*time.Second)
	}

	return
}

func GetPermanentStrQrImg(sceneStr string) (imgUrl string) {
	cache := help.Cache
	t := cache.Get("qr_img_" + sceneStr)
	if t != nil {
		imgUrl = string(t.([]uint8))
		return
	}

	expire := 30 * 24 * 3600
	m := map[string]interface{}{}
	m["action_name"] = "QR_LIMIT_STR_SCENE"
	m["action_info"] = map[string]interface{}{"scene": map[string]string{"scene_str": sceneStr}}
	m["expire_seconds"] = expire

	imgUrl = GetQrCodeImg(m)
	if len(imgUrl) > 0 {
		cache.Put("qr_img_"+sceneStr, imgUrl, time.Duration(expire)*time.Second)
	}

	return
}

func GetShortUrl(urlLong string) (urlShort string) {
	m := map[string]string{}
	m["action"] = "long2short"
	m["long_url"] = urlLong

	url := ApiUrl + "/cgi-bin/shorturl?access_token=" + GetAccessToken()
	req := httplib.Post(url)
	req.JSONBody(m)
	data := make(map[string]interface{})
	req.ToJSON(&data)

	if v, ok := data["short_url"]; ok {
		urlShort = v.(string)
	}

	return
}

func GetOpenId(c *context.Context) (openid string) {
	return wxpay.GetOpenId(c)
}

func CreateMenu(m map[string]interface{}) map[string]interface{} {
	url := ApiUrl + "/cgi-bin/menu/create?access_token=" + GetAccessToken()
	req := httplib.Post(url)
	req.JSONBody(m)
	data := make(map[string]interface{})
	req.ToJSON(&data)

	//fmt.Println(data)
	return data
}
