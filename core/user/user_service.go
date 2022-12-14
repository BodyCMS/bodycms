package user

import "github.com/BodyCMS/bodycms/lib/db"

type UserService struct {
	UserRepo *UserRepo
}

func (u *UserService) New() *UserService {
	return &UserService{
		UserRepo: &UserRepo{
			CmsDb: db.GetDbInstance(),
		},
	}
}

type UserWhere struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserService) FindOneId(id int, des *User) error {
	return u.UserRepo.FindOneId(id, des)
}

func (u *UserService) FindOne(where *UserWhere, des *User) error {
	return u.UserRepo.FindOne(where, des)
}

func (u *UserService) FindAll(where *UserWhere, des *[]User) error {
	return u.UserRepo.FindAll(where, des)
}

func (u *UserService) Create(des *User) error {
	return u.UserRepo.Create(des)
}

func (u *UserService) Update(where *UserWhere, data *UserUpdate, des *User) error {
	return u.UserRepo.Update(where, data, des)
}

func (u *UserService) Delete(where *UserWhere) error {
	return u.UserRepo.Delete(where)
}
