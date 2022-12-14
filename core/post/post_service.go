package post

import "github.com/BodyCMS/bodycms/lib/db"

type PostService struct {
	PostRepo *PostRepo
}

func (p *PostService) New() *PostService {
	return &PostService{
		PostRepo: &PostRepo{
			CmsDb: &db.CmsDb{},
		},
	}
}

type PostWhere struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type PostUpdate struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorId   int    `json:"author_id"`
	Author     string `json:"author"`
	Tags       string `json:"tags"`
	Categories string `json:"categories"`
}

func (p *PostService) FindOneId(id int, des *Post) error {
	return p.PostRepo.FindOneId(id, des)
}

func (p *PostService) FindOne(where *PostWhere, des *Post) error {
	return p.PostRepo.FindOne(where, des)
}

func (p *PostService) FindAll(where *PostWhere, des *[]Post) error {
	return p.PostRepo.FindAll(where, des)
}

func (p *PostService) Create(des *Post) error {
	return p.PostRepo.Create(des)
}

func (p *PostService) Update(where *PostWhere, data *PostUpdate, des *Post) error {
	return p.PostRepo.Update(where, data, des)
}

func (p *PostService) Delete(where *PostWhere) error {
	return p.PostRepo.Delete(where)
}
