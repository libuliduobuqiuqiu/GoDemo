package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormdemo/model"

	"github.com/go-faker/faker/v4"
)

func InnsertCompanyRows() error {
	db := gormdemo.GetDB()

	var companys []*model.Company
	for i := 0; i < 10; i++ {
		tmpCompany := model.Company{}
		err := faker.FakeData(&tmpCompany)
		if err != nil {
			return err
		}
		companys = append(companys, &tmpCompany)
	}

	result := db.Create(companys)
	if result.Error != nil {
		return result.Error
	}

	gormdemo.PrintRecords(companys, result)
	return nil
}

func SelectCompanyRows() {
	var codes []*model.Company
	db := gormdemo.GetDB()

	result := db.Model(&model.Company{}).Find(&codes)
	fmt.Println(codes)
	fmt.Println(result.Error, result.RowsAffected)

	var users []*model.User
	db.Model(&model.User{}).Find(&users)

	for k, user := range users {
		user.CompanyID = codes[k%len(codes)].Code

	}

	for _, tmpUser := range users {
		fmt.Println(tmpUser)
		result = db.Model(&model.User{}).Where("id = ?", tmpUser.ID).Updates(tmpUser)
		fmt.Println(result.Error, result.RowsAffected)
	}
}
