package services

import (
	"gorm.io/gorm"
	"my_gql_server/graph/model"
)

type Services interface {
	UserService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type UserService interface {
	GetUser(id uint) (*model.User, error)
}

type services struct {
	*userService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

func New(db *gorm.DB) Services {
	return &services{
		userService: &userService{db: db},
	}
}
