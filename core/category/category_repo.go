package category

import "github.com/BodyCMS/bodycms/lib/db"

type CategoryRepo struct {
	CmsDb *db.CmsDb
}

func (c *CategoryRepo) SetEngine(engine *db.CmsDb) error {
	c.CmsDb = engine
	return nil
}

func (c *CategoryRepo) AutoMigration() error {
	return c.CmsDb.Pdb.AutoMigrate(&Category{})
}

func (c *CategoryRepo) FindOneId(id int, des interface{}) error {
	return c.CmsDb.Pdb.First(des, id).Error
}

func (c *CategoryRepo) FindOne(where interface{}, des interface{}) error {
	return c.CmsDb.Pdb.Where(where).First(des).Error
}

func (c *CategoryRepo) FindAll(where interface{}, des interface{}) error {
	return c.CmsDb.Pdb.Where(where).Find(des).Error
}

func (c *CategoryRepo) Create(des interface{}) error {
	return c.CmsDb.Pdb.Create(des).Error
}

func (c *CategoryRepo) Update(where interface{}, data interface{}, des interface{}) error {
	return c.CmsDb.Pdb.Model(&des).Where(where).Updates(data).Error
}

func (c *CategoryRepo) Delete(where interface{}) error {
	return c.CmsDb.Pdb.Where(where).Delete(&Category{}).Error
}
