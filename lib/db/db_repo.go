package db

type IDbRepo interface {
	SetEngine(engine *CmsDb) error
	AutoMigration() error
}
