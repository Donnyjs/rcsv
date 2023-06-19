package repo

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/entity"
)

type InscriptionRepository interface {
	Insert(inscription *po.Inscription) (err error)
	List(w *entity.MysqlWhere) (list []*po.Inscription, err error)
	CheckInscription(w *entity.MysqlWhere) (inscription *po.Inscription, err error)
}

type inscriptionRepository struct {
}

func NewInscriptionRepository() InscriptionRepository {
	return &inscriptionRepository{}
}

func (r *inscriptionRepository) Insert(inscription *po.Inscription) (err error) {
	db := xmysql.GetDB()
	err = db.Create(inscription).Error
	return
}

func (r *inscriptionRepository) CheckInscription(w *entity.MysqlWhere) (inscription *po.Inscription, err error) {
	inscription = new(po.Inscription)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(inscription).Error
	return
}

func (r *inscriptionRepository) List(w *entity.MysqlWhere) (list []*po.Inscription, err error) {
	list = make([]*po.Inscription, 0)
	db := xmysql.GetDB()
	err = db.Model(po.Inscription{}).
		Select("id,inscription,inscription_id").
		Where(w.Query, w.Args...).
		Limit(w.Limit).
		Find(&list).Error
	return
}
