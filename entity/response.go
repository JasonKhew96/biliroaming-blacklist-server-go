package entity

type RespStatusCompatData struct {
	UID         int64  `json:"uid"`
	IsBlacklist bool   `json:"is_blacklist"`
	IsWhitelist bool   `json:"is_whitelist"`
	Reason      string `json:"reason"`
}

type RespStatusCompat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    RespStatusCompatData `json:"data"`
}

type RespStatusData struct {
	UID          int64 `json:"uid"`
	IsBlacklist  bool  `json:"is_blacklist"`
	IsWhitelist  bool  `json:"is_whitelist"`
	Status       int8  `json:"status"`
	BanUntil     int64 `json:"ban_until"`
	Counter      int64 `json:"counter,omitempty"`
	LastAccessed int64 `json:"last_accessed,omitempty"`
}

type RespStatus struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    *RespStatusData `json:"data,omitempty"`
}
