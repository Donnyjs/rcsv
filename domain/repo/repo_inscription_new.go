package repo

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xmysql"
)

type InscriptionNewRepository interface {
	Insert(inscription *po.InscriptionNew) (err error)
}

type inscriptionNewRepository struct {
}

func NewInscriptionNewRepository() InscriptionNewRepository {
	return &inscriptionNewRepository{}
}

func (r *inscriptionNewRepository) Insert(inscription *po.InscriptionNew) (err error) {
	db := xmysql.GetDB()
	err = db.Create(inscription).Error
	return
}
