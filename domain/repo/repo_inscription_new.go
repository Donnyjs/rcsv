package repo

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/entity"
)

type InscriptionNewRepository interface {
	Insert(inscription *po.InscriptionNew) (err error)
	List(w *entity.MysqlWhere) (list []po.InscriptionNew, err error)
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

func (r *inscriptionNewRepository) List(w *entity.MysqlWhere) (list []po.InscriptionNew, err error) {
	list = make([]po.InscriptionNew, 0)
	db := xmysql.GetDB()
	err = db.Model(po.Inscription{}).
		Select("id,inscription,inscription_id, recursive_num, genesis_block_height, genesis_timestamp, content_length").
		Where(w.Query, w.Args...).
		Offset(w.Offset).
		Order(w.Sort).
		Limit(w.Limit).
		Find(&list).Error
	return
}
