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

type CardByMidCard struct {
	Mid  string `json:"mid"`
	Name string `json:"name"`
}

type CardByMid struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Card    *CardByMidCard `json:"card,omitempty"`
}

type Nav struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		WbiImg struct {
			ImgUrl string `json:"img_url"`
			SubUrl string `json:"sub_url"`
		} `json:"wbi_img"`
	} `json:"data"`
}
