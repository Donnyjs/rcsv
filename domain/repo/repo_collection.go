package repo

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xmysql"
	"strconv"
)

type RCSVRepository interface {
	Insert(inscription *po.RcsvCollection) (err error)
	List(sort string, page, limit int) (list []po.InscriptionInfoWithMetaNum, err error)
	CountCollection() (count int64, err error)
	ListWithNum(numStart, numEnd int64) (list []po.InscriptionInfoWithMetaNum, err error)
}

type rcsvRepository struct {
}

func NewRCSVRepository() RCSVRepository {
	return &rcsvRepository{}
}

func (rcsvRepository) Insert(inscription *po.RcsvCollection) (err error) {
	db := xmysql.GetDB()
	err = db.Create(inscription).Error
	return
}

func (rcsvRepository) List(sort string, page, limit int) (list []po.InscriptionInfoWithMetaNum, err error) {
	list = make([]po.InscriptionInfoWithMetaNum, 0)
	db := xmysql.GetDB()
	db.Table("inscription_info").
		Select("inscription_info.*, rcsv_collection.meta_num").
		Joins("INNER JOIN rcsv_collection ON inscription_info.inscription_id = rcsv_collection.inscription_id").
		Order(sort).
		Limit(limit).Offset(limit * (page - 1)).
		Find(&list)
	return
}

func (rcsvRepository) CountCollection() (count int64, err error) {
	db := xmysql.GetDB()
	err = db.Model(po.RcsvCollection{}).Count(&count).Error
	return
}

func (rcsvRepository) ListWithNum(numStart, numEnd int64) (list []po.InscriptionInfoWithMetaNum, err error) {

	list = make([]po.InscriptionInfoWithMetaNum, 0)
	db := xmysql.GetDB()
	db.Table("inscription_info").
		Select("inscription_info.*, rcsv_collection.meta_num").
		Joins("INNER JOIN rcsv_collection ON inscription_info.inscription_id = rcsv_collection.inscription_id").
		Where("rcsv_collection.meta_num >=" + strconv.FormatInt(numStart, 10) + " and rcsv_collection.meta_num<=" + strconv.FormatInt(numEnd, 10)).
		Find(&list)
	return
}
