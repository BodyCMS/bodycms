package post

import (
	"github.com/BodyCMS/bodycms/lib/db"
)

type PostRepo struct {
	CmsDb *db.CmsDb
}

func (p *PostRepo) SetEngine(engine *db.CmsDb) error {
	p.CmsDb = engine
	return nil
}

func (p *PostRepo) AutoMigration() error {
	return p.CmsDb.Pdb.AutoMigrate(&Post{})
}

func (p *PostRepo) FindOneId(id int, des interface{}) error {
	return p.CmsDb.Pdb.First(des, id).Error
}

func (p *PostRepo) FindOne(where interface{}, des interface{}) error {
	return p.CmsDb.Pdb.Where(where).First(des).Error
}

func (p *PostRepo) FindAll(where interface{}, des interface{}) error {
	return p.CmsDb.Pdb.Where(where).Find(des).Error
}

func (p *PostRepo) Create(des interface{}) error {
	return p.CmsDb.Pdb.Create(des).Error
}

func (p *PostRepo) Update(where interface{}, data interface{}, des interface{}) error {
	return p.CmsDb.Pdb.Model(&des).Where(where).Updates(data).Error
}

func (p *PostRepo) Delete(where interface{}) error {
	return p.CmsDb.Pdb.Where(where).Delete(&Post{}).Error
}
