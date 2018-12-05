package huobi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"net/url"
)


// HMAC SHA256加密
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

// 对map进行排序
func MapSortByKey(mapValue map[string]string) map[string]string {
	var keys []string
	for key := range mapValue {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	mapReturn := make(map[string]string)
	for _, key := range keys {
		mapReturn[key] = mapValue[key]
	}
	return mapReturn
}

// 将map格式请求参数装换为字符串格式，并按照map的key升序排列
func MapUrlQueryBySort(mapParams map[string]string) string {
	var keys []string
	for key := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var strParams string
	for _, key := range keys {
		strParams += key + "=" + mapParams[key] + "&"
	}
	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}
	return strParams
}

// 创建签名
func CreateSign(mapParams map[string]string, strMethod, strHostUrl, strRequestPath, strSecretKey string) (string, error) {
	mapCloned := make(map[string]string)
	for key, value := range mapParams {
		mapCloned[key] = url.QueryEscape(value)
	}
	strParams := MapUrlQueryBySort(mapCloned)

	strPayload := strMethod + "\n" + strHostUrl + "\n" + strRequestPath + "\n" + strParams
	return HmacSha256Base64Signer(strPayload, strSecretKey)
}
