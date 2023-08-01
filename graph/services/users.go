package services

import (
	"gorm.io/gorm"
	"my_gql_server/graph/db"
	"my_gql_server/graph/model"
	"strconv"
)

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Work:  user.Work,
	}
}

type userService struct {
	db *gorm.DB
}

func (u userService) GetUser(id uint) (*model.User, error) {
	var user *db.User
	result := u.db.First(&user, id)
	return convertUser(user), result.Error
}
