package tag

import (
	"github.com/BodyCMS/bodycms/lib/db"
)

type TagRepo struct {
	CmsDb *db.CmsDb
}

func (t *TagRepo) SetEngine(engine *db.CmsDb) error {
	t.CmsDb = engine
	return nil
}

func (t *TagRepo) AutoMigration() error {
	return t.CmsDb.Pdb.AutoMigrate(&Tag{})
}
