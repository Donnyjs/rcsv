package po

import "rcsv/pkg/entity"

type Inscription struct {
	entity.GormEntityTs
	Id            string `gorm:"column:id;primary_key" json:"id"`
	Inscription   int64  `gorm:"column:inscription;NOT NULL" json:"inscription"`
	InscriptionId string `gorm:"column:inscription_id;NOT NULL" json:"inscription_id"`
}

func (Inscription) TableName() string {
	return "inscription_info"
}