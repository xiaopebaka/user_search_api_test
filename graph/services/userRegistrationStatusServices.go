package services

import (
	"gorm.io/gorm"
	"my_gql_server/graph/db"
	"my_gql_server/graph/model"
	"strconv"
)

func convertUserRegistrationStatus(userRegistrationStatus *db.UserRegistrationStatus) *model.UserRegistrationStatus {
	return &model.UserRegistrationStatus{
		ID:              strconv.Itoa(int(userRegistrationStatus.ID)),
		AccountID:       userRegistrationStatus.AccountID,
		PreRegisteredAt: userRegistrationStatus.PreRegisteredAt,
		RegisteredAt:    userRegistrationStatus.RegisteredAt,
	}
}

type userRegistrationStatusService struct {
	db *gorm.DB
}

func (u userRegistrationStatusService) GetUserRegistrationStatus(accountId string) (*model.UserRegistrationStatus, error) {
	var userRegistrationStatus *db.UserRegistrationStatus
	result := u.db.Where("account_id = ?", accountId).First(&userRegistrationStatus)
	return convertUserRegistrationStatus(userRegistrationStatus), result.Error
}
