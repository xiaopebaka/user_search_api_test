package services

import (
	"gorm.io/gorm"
	"my_gql_server/graph/db"
	"my_gql_server/graph/model"
	"strconv"
)

func convertUser(user *db.Users) *model.User {
	return &model.User{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Name,
		Birth: user.Birth,
	}
}

type userService struct {
	db *gorm.DB
}

func (u userService) GetUser(id uint) (*model.User, error) {
	var user *db.Users
	result := u.db.First(&user, id)
	return convertUser(user), result.Error
}
