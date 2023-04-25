package utils

import (
	"biliroaming-blacklist-server-go/entity"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

	rawByte, err := ioutil.ReadAll(resp.Body)
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

	rawByte, err := ioutil.ReadAll(resp.Body)
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

	rawByte, err := ioutil.ReadAll(resp.Body)
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
