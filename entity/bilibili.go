package entity

type SpaceAccInfoData struct {
	Mid  int64  `json:"mid"`
	Name string `json:"name"`
}

type SpaceAccInfo struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    *SpaceAccInfoData `json:"data,omitempty"`
}

type CardByMid struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Card    *SpaceAccInfoData `json:"card,omitempty"`
}
