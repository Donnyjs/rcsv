package dto_inscription

type InscriptionListResp struct {
	Inscription   int    `json:"inscription"`
	InscriptionId string `json:"inscription_id"`
}

type InscriptionListReq struct {
	Type  string `form:"type"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}
