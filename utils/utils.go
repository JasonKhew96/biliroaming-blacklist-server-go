package utils

import (
	"biliroaming-blacklist-server-go/entity"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const USER_AGENT = "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0"

func GetUserInfo(uid int64) (*entity.SpaceAccInfoData, error) {
	data, err := GetBiliAccInfo(uid)
	if err != nil {
		log.Println(err)
	} else {
		return data, nil
	}
	return GetCardByMid(uid)
}

func GetBiliAccInfo(uid int64) (*entity.SpaceAccInfoData, error) {
	reqUrl := fmt.Sprintf("https://api.bilibili.com/x/space/acc/info?mid=%d", uid)
	signedReqUrl, err := signAndGenerateURL(reqUrl)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, signedReqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", USER_AGENT)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("content type: %s", resp.Header.Get("Content-Type"))
	}

	rawByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data entity.SpaceAccInfo
	err = json.Unmarshal(rawByte, &data)
	if err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", data.Code, data.Message)
	}

	if data.Data == nil {
		return nil, fmt.Errorf("data is nil")
	}

	if data.Data.Mid != uid {
		return nil, fmt.Errorf("uid not match")
	}

	return data.Data, nil
}

func GetCardByMid(uid int64) (*entity.SpaceAccInfoData, error) {
	reqUrl := fmt.Sprintf("https://account.bilibili.com/api/member/getCardByMid?mid=%d", uid)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", USER_AGENT)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	// if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
	// 	return nil, fmt.Errorf("content type: %s", resp.Header.Get("Content-Type"))
	// }

	rawByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data entity.CardByMid
	err = json.Unmarshal(rawByte, &data)
	if err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", data.Code, data.Message)
	}

	if data.Card == nil {
		return nil, fmt.Errorf("data is nil")
	}

	mid, err := strconv.ParseInt(data.Card.Mid, 10, 64)
	if err != nil {
		return nil, err
	}

	if mid != uid {
		return nil, fmt.Errorf("uid not match")
	}

	return &entity.SpaceAccInfoData{
		Mid:  mid,
		Name: data.Card.Name,
	}, nil
}

func GetMyInfo(key string) (*entity.SpaceAccInfoData, error) {
	values := url.Values{}
	values.Set("access_key", key)

	params, err := signParams(values)
	if err != nil {
		return nil, err
	}

	reqUrl := fmt.Sprintf("https://api.bilibili.com/x/space/myinfo?%s", params)

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", USER_AGENT)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("content type: %s", resp.Header.Get("Content-Type"))
	}

	rawByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data entity.SpaceAccInfo
	err = json.Unmarshal(rawByte, &data)
	if err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", data.Code, data.Message)
	}

	if data.Data == nil {
		return nil, fmt.Errorf("data is nil")
	}

	return data.Data, nil
}

func signParams(values url.Values) (string, error) {
	sign, err := getSign(values)
	if err != nil {
		return "", err
	}
	values.Set("sign", sign)
	return values.Encode(), nil
}

func getSign(values url.Values) (string, error) {
	appkey := "1d8b6e7d45233436"
	appsec := "560c52ccd288fed045859ed18bffd973"

	values.Set("ts", strconv.FormatInt(time.Now().Unix(), 10))
	values.Set("appkey", appkey)

	encoded := values.Encode() + appsec
	data := []byte(encoded)
	return fmt.Sprintf("%x", md5.Sum(data)), nil
}

func ParseDuration(duration string) (*time.Time, error) {
	if len(duration) <= 1 {
		return nil, fmt.Errorf("duration too short")
	}

	now := time.Now()

	addStr := duration[:len(duration)-1]
	add, err := strconv.Atoi(addStr)
	if err != nil {
		return nil, err
	}

	unit := duration[len(duration)-1:]

	switch unit {
	case "h", "H": // hour
		newTime := now.Add(time.Hour * time.Duration(add))
		return &newTime, nil
	case "d", "D": // day
		newTime := now.AddDate(0, 0, add)
		return &newTime, nil
	case "w", "W": // week
		newTime := now.AddDate(0, 0, add*7)
		return &newTime, nil
	case "m", "M": // month
		newTime := now.AddDate(0, add, 0)
		return &newTime, nil
	case "y", "Y": // year
		newTime := now.AddDate(add, 0, 0)
		return &newTime, nil
	default:
		return nil, fmt.Errorf("invalid duration unit: %s", unit)
	}
}

func ConvertNextLine(s string) template.HTML {
	return template.HTML(strings.ReplaceAll(s, "\n", "<br>"))
}

var (
	mixinKeyEncTab = []int{
		46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
		33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
		61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
		36, 20, 34, 44, 52,
	}
	cache          sync.Map
	lastUpdateTime time.Time
)

func signAndGenerateURL(urlStr string) (string, error) {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	imgKey, subKey, err := getWbiKeysCached()
	if err != nil {
		return "", err
	}
	query := urlObj.Query()
	params := map[string]string{}
	for k, v := range query {
		params[k] = v[0]
	}
	newParams := encWbi(params, imgKey, subKey)
	for k, v := range newParams {
		query.Set(k, v)
	}
	urlObj.RawQuery = query.Encode()
	newUrlStr := urlObj.String()
	return newUrlStr, nil
}

func encWbi(params map[string]string, imgKey, subKey string) map[string]string {
	mixinKey := getMixinKey(imgKey + subKey)
	currTime := strconv.FormatInt(time.Now().Unix(), 10)
	params["wts"] = currTime

	// Sort keys
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Remove unwanted characters
	for k, v := range params {
		v = sanitizeString(v)
		params[k] = v
	}

	// Build URL parameters
	query := url.Values{}
	for _, k := range keys {
		query.Set(k, params[k])
	}
	queryStr := query.Encode()

	// Calculate w_rid
	hash := md5.Sum([]byte(queryStr + mixinKey))
	params["w_rid"] = hex.EncodeToString(hash[:])
	return params
}

func getMixinKey(orig string) string {
	var str strings.Builder
	for _, v := range mixinKeyEncTab {
		if v < len(orig) {
			str.WriteByte(orig[v])
		}
	}
	return str.String()[:32]
}

func sanitizeString(s string) string {
	unwantedChars := []string{"!", "'", "(", ")", "*"}
	for _, char := range unwantedChars {
		s = strings.ReplaceAll(s, char, "")
	}
	return s
}

func updateCache() error {
	if time.Since(lastUpdateTime).Minutes() < 10 {
		return nil
	}
	imgKey, subKey, err := getWbiKeys()
	if err != nil {
		return err
	}
	cache.Store("imgKey", imgKey)
	cache.Store("subKey", subKey)
	lastUpdateTime = time.Now()
	return nil
}

func getWbiKeysCached() (string, string, error) {
	if err := updateCache(); err != nil {
		return "", "", err
	}
	imgKeyI, ok := cache.Load("imgKey")
	if !ok {
		return "", "", errors.New("imgKey not found")
	}
	subKeyI, ok := cache.Load("subKey")
	if !ok {
		return "", "", errors.New("subKey not found")
	}
	return imgKeyI.(string), subKeyI.(string), nil
}

func getWbiKeys() (string, string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.bilibili.com/x/web-interface/nav", nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", USER_AGENT)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var data entity.Nav
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", err
	}
	imgURL := data.Data.WbiImg.ImgUrl
	subURL := data.Data.WbiImg.SubUrl
	imgKey := strings.Split(strings.Split(imgURL, "/")[len(strings.Split(imgURL, "/"))-1], ".")[0]
	subKey := strings.Split(strings.Split(subURL, "/")[len(strings.Split(subURL, "/"))-1], ".")[0]
	return imgKey, subKey, nil
}
