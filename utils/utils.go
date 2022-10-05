package utils

import (
	"biliroaming-blacklist-server-go/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"

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

	var data entity.SpaceAccInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
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
