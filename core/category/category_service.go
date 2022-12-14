package category

import "github.com/BodyCMS/bodycms/lib/db"

type CategoryService struct {
	CategoryRepo *CategoryRepo
}

func (c *CategoryService) New() *CategoryService {
	return &CategoryService{
		CategoryRepo: &CategoryRepo{
			CmsDb: db.GetDbInstance(),
		},
	}
}

type CategoryWhere struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryUpdate struct {
	Name string `json:"name"`
}

func (c *CategoryService) FindOneId(id int, des *Category) error {
	return c.CategoryRepo.FindOneId(id, des)
}

func (c *CategoryService) FindOne(where *CategoryWhere, des *Category) error {
	return c.CategoryRepo.FindOne(where, des)
}

func (c *CategoryService) FindAll(where *CategoryWhere, des *[]Category) error {
	return c.CategoryRepo.FindAll(where, des)
}

func (c *CategoryService) Create(des *Category) error {
	return c.CategoryRepo.Create(des)
}

func (c *CategoryService) Update(where *CategoryWhere, data *CategoryUpdate, des *Category) error {
	return c.CategoryRepo.Update(where, data, des)
}

func (c *CategoryService) Delete(where *CategoryWhere) error {
	return c.CategoryRepo.Delete(where)
}
