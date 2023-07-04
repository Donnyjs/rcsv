package repo

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xmysql"
)

type ParentSonRepository interface {
	Insert(inscription *po.ParentSon) (err error)
	BatchInsert(insList []*po.ParentSon, batchSize int) (err error)
}

type parentSonRepository struct {
}

func NewParentSonRepository() ParentSonRepository {
	return &parentSonRepository{}
}

func (r *parentSonRepository) Insert(inscription *po.ParentSon) (err error) {
	db := xmysql.GetDB()
	err = db.Create(inscription).Error
	return
}

func (r *parentSonRepository) BatchInsert(insList []*po.ParentSon, batchSize int) (err error) {
	db := xmysql.GetDB()
	err = db.CreateInBatches(insList, batchSize).Error
	return
}
