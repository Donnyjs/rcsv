package svc_inscription_monitor

type HiroInscriptionResp struct {
	Limit   int64    `json:"limit"`
	Offset  int64    `json:"offset"`
	Total   int64    `json:"total"`
	Results []Result `json:"results"`
}

type Result struct {
	Id                 string `json:"id"`
	Number             int64  `json:"number"`
	GenesisBlockHeight int64  `json:"genesis_block_height""`
	//还有一些参数，目前用不上
}
