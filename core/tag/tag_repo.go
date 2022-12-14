package tag

import (
	"github.com/BodyCMS/bodycms/lib/db"
)

// Tag Repository implementation of IDbRepo
type TagRepo struct {
	CmsDb *db.CmsDb
}

type TagWhere struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagUpdate struct {
	Name string `json:"name"`
}

func (t *TagRepo) SetEngine(engine *db.CmsDb) error {
	t.CmsDb = engine
	return nil
}

func (t *TagRepo) AutoMigration() error {
	return t.CmsDb.Pdb.AutoMigrate(&Tag{})
}

func (t *TagRepo) FindOneId(id int, des interface{}) error {
	return t.CmsDb.Pdb.First(des, id).Error
}

func (t *TagRepo) FindOne(where interface{}, des interface{}) error {
	return t.CmsDb.Pdb.Where(where).First(des).Error
}

func (t *TagRepo) FindAll(where interface{}, des interface{}) error {
	return t.CmsDb.Pdb.Where(where).Find(des).Error
}

func (t *TagRepo) Create(des interface{}) error {
	return t.CmsDb.Pdb.Create(des).Error
}

func (t *TagRepo) Update(where interface{}, data interface{}, des interface{}) error {
	return t.CmsDb.Pdb.Model(&des).Where(where).Updates(data).Error
}

func (t *TagRepo) Delete(where interface{}) error {
	return t.CmsDb.Pdb.Where(where).Delete(&Tag{}).Error
}
