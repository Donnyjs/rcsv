package repo

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/constant"
	"rcsv/pkg/entity"
)

type InscriptionRepository interface {
	Insert(inscription *po.Inscription) (err error)
	Update(w *entity.MysqlUpdate) error
	List(w *entity.MysqlWhere) (list []po.Inscription, err error)
	CountDoodinals() (count int64, err error)
	CountOthers() (count int64, err error)
	CheckInscription(w *entity.MysqlWhere) (inscription *po.Inscription, err error)
	QueryByNumber(w *entity.MysqlWhere) (inscription *po.Inscription, err error)
}

type inscriptionRepository struct {
}

func NewInscriptionRepository() InscriptionRepository {
	return &inscriptionRepository{}
}

func (r *inscriptionRepository) QueryByNumber(w *entity.MysqlWhere) (inscription *po.Inscription, err error) {
	inscription = new(po.Inscription)
	db := xmysql.GetDB()
	err = db.Select("inscription,inscription_id").Where(w.Query, w.Args...).Find(inscription).Error
	return
}

func (r *inscriptionRepository) Insert(inscription *po.Inscription) (err error) {
	db := xmysql.GetDB()
	err = db.Create(inscription).Error
	return
}

func (r *inscriptionRepository) Update(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(po.Inscription{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *inscriptionRepository) CheckInscription(w *entity.MysqlWhere) (inscription *po.Inscription, err error) {
	inscription = new(po.Inscription)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(inscription).Error
	return
}

func (r *inscriptionRepository) List(w *entity.MysqlWhere) (list []po.Inscription, err error) {
	list = make([]po.Inscription, 0)
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

func (r *inscriptionRepository) CountDoodinals() (count int64, err error) {
	db := xmysql.GetDB()
	err = db.Model(po.Inscription{}).Where("data_type = ?", constant.DOODINALS).Count(&count).Error
	return
}

func (r *inscriptionRepository) CountOthers() (count int64, err error) {
	db := xmysql.GetDB()
	err = db.Model(po.Inscription{}).Where("data_type = ?", constant.DATA_RCSV_IO).Count(&count).Error
	return
}
