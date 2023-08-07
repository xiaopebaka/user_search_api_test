package static

import (
	"fmt"
	"gorm.io/gorm"
	"my_gql_server/graph/model"
)

func createLocations(db *gorm.DB) {
	var location model.Location
	if err := db.Where("name = ?", "default").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// default user not found, create
			defaultUser := User{Name: "default"}
			if err := db.Create(&defaultUser).Error; err != nil {
				// handle error
				fmt.Println(err)
			}
		} else {
			// handle error
			fmt.Println(err)
		}
	}
	if err := db.Model(&location{})
}
