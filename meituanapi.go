package meituanapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/guonaihong/gout"
)

type Api struct {
	appKey    string
	appSecret string
}

func NewTbkApi(appKey, appSecret string) *Api {
	return &Api{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

// OrderList 订单列表查询
func (a *Api) OrderList(param OrderListParams) (*OrderListResp, error) {
	p, err := a.parameterCombination(param)
	if err != nil {
		return nil, err
	}
	resp := OrderListResp{}
	err = gout.GET("https://runion.meituan.com/api/orderList").BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// OrderInfo 订单详情查询
func (a *Api) OrderInfo(param OrderInfoParams) (*OrderInfoResp, error) {
	p, err := a.parameterCombination(param)
	if err != nil {
		return nil, err
	}
	resp := OrderInfoResp{}
	err = gout.GET("https://runion.meituan.com/api/rtnotify").Debug(true).BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CouponList 优惠券列表查询
func (a *Api) CouponList(param CouponListParams) (*CouponListResp, error) {
	p, err := a.parameterCombination(param)
	if err != nil {
		return nil, err
	}
	resp := CouponListResp{}
	err = gout.GET("https://runion.meituan.com/api/couponList").BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateLink 自助取链
func (a *Api) GenerateLink(param GenerateLinkParams) (*GenerateLinkResp, error) {
	p, err := a.parameterCombination(param)
	if err != nil {
		return nil, err
	}
	resp := GenerateLinkResp{}
	err = gout.GET("https://runion.meituan.com/generateLink").BindJSON(&resp).SetQuery(p).Do()
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *Api) parameterCombination(param interface{}) (map[string]interface{}, error) {
	p := a.getPublicParameters()
	_p, err := Struct2Map(param)
	if err != nil {
		return nil, err
	}

	for k, v := range _p {
		if _, ok := v.(map[string]interface{}); ok {
			dataType, _ := Marshal(v)
			dataString := string(dataType)
			p[k] = dataString
		} else if _, ok := v.([]interface{}); ok {
			dataType, _ := Marshal(v)
			dataString := string(dataType)
			p[k] = dataString
		} else {
			p[k] = v
		}
	}
	sign := a.getSign(p)
	p["sign"] = sign
	return p, nil
}

func (a *Api) getSign(param map[string]interface{}) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := a.appSecret
	for _, key := range keys {
		str += key
		str += fmt.Sprint(param[key])
	}
	str += a.appSecret
	return GetMd5Encode(str)
}

func (a *Api) getPublicParameters() map[string]interface{} {
	return map[string]interface{}{
		"key": a.appKey,
		"ts":  time.Now().Unix(),
	}
}

func GetMd5Encode(data string) string {
	h := md5.New()
	_, _ = h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	j, _ := json.Marshal(obj)
	err := json.Unmarshal(j, &data)
	return data, err
}