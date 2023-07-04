package po

import "rcsv/pkg/entity"

type ParentSon struct {
	entity.GormEntityTs
	Id          int16  `gorm:"column:id;primary_key" json:"id"`
	InsUuid     string `gorm:"column:ins_uuid;NOT NULL" json:"ins_uuid"`
	InsId       string `gorm:"column:ins_id;NOT NULL" json:"ins_id"`
	Count       int    `gorm:"default:0;column:count;NOT NULL" json:"count"`
	DupCount    int64  `gorm:"default:0;column:dup_count;NOT NULL" json:"dup_count"`
	ParentInsId string `gorm:"column:parent_ins_id;NOT NULL" json:"parent_ins_id"`
}

func (ParentSon) TableName() string {
	return "parent_son"
}
