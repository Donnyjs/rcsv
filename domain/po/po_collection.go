package po

import "rcsv/pkg/entity"

type RcsvCollection struct {
	entity.GormEntityTs
	Id            string `gorm:"column:id;primary_key" json:"id"`
	InscriptionId string `gorm:"column:inscription_id;NOT NULL" json:"inscription_id"`
	MetaName      string `gorm:"column:meta_name;NOT NULL" json:"meta_name"`
	MetaNum       int64  `gorm:"column:meta_num;NOT NULL" json:"meta_num"`
}

func (RcsvCollection) TableName() string {
	return "rcsv_collection"
}

// InscriptionInfoWithMetaNum 定义合并后的结构体
type InscriptionInfoWithMetaNum struct {
	Inscription
	MetaNum int64 `gorm:"column:meta_num;NOT NULL" json:"meta_num"`
}
