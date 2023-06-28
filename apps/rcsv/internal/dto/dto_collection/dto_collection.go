package dto_collection

import "rcsv/domain/po"

type CollectionListResp struct {
	DoodinalsTotal int64                           `json:"collection_total"`
	Data           []po.InscriptionInfoWithMetaNum `json:"data"`
}

type InscriptionData struct {
	Inscription      int    `json:"inscription"`
	InscriptionId    string `json:"inscription_id"`
	ContentLength    int64  `json:"content_length"`
	RecursiveNum     int64  `json:"recursive_num"`
	GenesisTimestamp int64  `json:"genesis_timestamp"`
}

type CollectionListReq struct {
	Sort  string `form:"sort"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}
