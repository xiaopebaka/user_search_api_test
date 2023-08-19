package services

import (
	"gorm.io/gorm"
	"my_gql_server/graph/model"
)

type Services interface {
	UserService
	UserRegistrationStatusService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type UserService interface {
	GetUser(id uint) (*model.User, error)
}
type UserRegistrationStatusService interface {
	GetUserRegistrationStatus(accountId string) (*model.UserRegistrationStatus, error)
}

type services struct {
	*userService
	*userRegistrationStatusService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

func New(db *gorm.DB) Services {
	return &services{
		userService:                   &userService{db: db},
		userRegistrationStatusService: &userRegistrationStatusService{db: db},
	}
}
