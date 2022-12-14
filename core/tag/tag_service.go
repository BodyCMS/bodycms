package tag

import "github.com/BodyCMS/bodycms/lib/db"

type TagService struct {
	TagRepo *TagRepo
}

func (t *TagService) New() *TagService {
	return &TagService{
		TagRepo: &TagRepo{
			CmsDb: db.GetDbInstance(),
		},
	}
}

func (t *TagService) FindOneId(id int, des *Tag) error {
	return t.TagRepo.FindOneId(id, des)
}

func (t *TagService) FindOne(where *TagWhere, des *Tag) error {
	return t.TagRepo.FindOne(where, des)
}

func (t *TagService) FindAll(where *TagWhere, des *[]Tag) error {
	return t.TagRepo.FindAll(where, des)
}

func (t *TagService) Create(des *Tag) error {
	return t.TagRepo.Create(des)
}

func (t *TagService) Update(where *TagWhere, data *TagUpdate, des *Tag) error {
	return t.TagRepo.Update(where, data, des)
}

func (t *TagService) Delete(where *TagWhere) error {
	return t.TagRepo.Delete(where)
}
