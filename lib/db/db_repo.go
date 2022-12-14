package db

type IDbRepo interface {
	SetEngine(engine *CmsDb) error
	AutoMigration() error
	FindOneId(id int, des interface{}) error
	FindOne(where interface{}, des interface{}) error
	FindAll(where interface{}, des interface{}) error
	Create(des interface{}) error
	Update(where interface{}, data interface{}, des interface{}) error
	Delete(where interface{}) error
}
