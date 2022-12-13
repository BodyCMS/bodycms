package tag

import (
	"github.com/BodyCMS/bodycms/lib/db"
)

type TagRepo struct {
	cmsDb *db.CmsDb
}

func (t *TagRepo) SetEngine(engine *db.CmsDb) error {
	t.cmsDb = engine
	return nil
}

func (t *TagRepo) AutoMigration() error {
	return t.cmsDb.Pdb.AutoMigrate(&Tag{})
}
