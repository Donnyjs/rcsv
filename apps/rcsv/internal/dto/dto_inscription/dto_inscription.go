package dto_inscription

type InscriptionListResp struct {
	DoodinalsTotal int               `json:"doodinals_total"`
	OthersTotal    int               `json:"others_total"`
	Data           []InscriptionData `json:"data"`
}

type InscriptionData struct {
	Inscription      int    `json:"inscription"`
	InscriptionId    string `json:"inscription_id"`
	ContentLength    int64  `json:"content_length"`
	RecursiveNum     int64  `json:"recursive_num"`
	GenesisTimestamp int64  `json:"genesis_timestamp"`
}

type InscriptionListReq struct {
	Sort  string `form:"sort"`
	Type  string `form:"type"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}
