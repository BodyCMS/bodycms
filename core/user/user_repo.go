package user

import "github.com/BodyCMS/bodycms/lib/db"

type UserRepo struct {
	CmsDb *db.CmsDb
}

func (u *UserRepo) SetEngine(engine *db.CmsDb) error {
	u.CmsDb = engine
	return nil
}

func (u *UserRepo) AutoMigration() error {
	return u.CmsDb.Pdb.AutoMigrate(&User{})
}

func (u *UserRepo) FindOneId(id int, des interface{}) error {
	return u.CmsDb.Pdb.First(des, id).Error
}

func (u *UserRepo) FindOne(where interface{}, des interface{}) error {
	return u.CmsDb.Pdb.Where(where).First(des).Error
}

func (u *UserRepo) FindAll(where interface{}, des interface{}) error {
	return u.CmsDb.Pdb.Where(where).Find(des).Error
}

func (u *UserRepo) Create(des interface{}) error {
	return u.CmsDb.Pdb.Create(des).Error
}

func (u *UserRepo) Update(where interface{}, data interface{}, des interface{}) error {
	return u.CmsDb.Pdb.Model(&des).Where(where).Updates(data).Error
}

func (u *UserRepo) Delete(where interface{}) error {
	return u.CmsDb.Pdb.Where(where).Delete(&User{}).Error
}
