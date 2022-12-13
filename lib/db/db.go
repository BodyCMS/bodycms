package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *CmsDb

type CmsDb struct {
	Options *DbOptions
	pdb     *gorm.DB
}

func init() {
	dbInstance = &CmsDb{
		Options: DefaultDbOptions().Assign(DbOptionsFromEnv()),
	}
}

func GetDbInstance() *CmsDb {
	return dbInstance
}

func (db *CmsDb) Connect() error {
	pdb, err := gorm.Open(
		postgres.Open(db.Options.GetConnectionString()),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	db.pdb = pdb

	return nil
}

func (db *CmsDb) AutoMigrate(model ...interface{}) error {
	for _, m := range model {
		err := db.pdb.AutoMigrate(m)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
