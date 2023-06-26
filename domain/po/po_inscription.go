package po

import "rcsv/pkg/entity"

type Inscription struct {
	entity.GormEntityTs
	Id                 string `gorm:"column:id;primary_key" json:"id"`
	Inscription        int64  `gorm:"column:inscription;NOT NULL" json:"inscription"`
	InscriptionId      string `gorm:"column:inscription_id;NOT NULL" json:"inscription_id"`
	DataType           string `gorm:"column:data_type;NOT NULL" json:"data_type"`
	RecursiveNum       int64  `gorm:"column:recursive_num;NOT NULL" json:"recursive_num"`
	GenesisBlockHeight int64  `gorm:"column:genesis_block_height;NOT NULL" json:"genesis_block_height"`
	GenesisTimestamp   int64  `gorm:"column:genesis_timestamp;NOT NULL" json:"genesis_timestamp"`
	ContentLength      int64  `gorm:"column:content_length;NOT NULL" json:"content_length"`
	Owner              string `gorm:"column:owner;NOT NULL" json:"owner"`
	Pic                string `gorm:"column:pic;NOT NULL" json:"pic"`
}

func (Inscription) TableName() string {
	return "inscription_info"
}
