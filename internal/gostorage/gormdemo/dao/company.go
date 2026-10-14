package dao

import (
	"errors"
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

func SelectCompanyRows() error {
	var codes []*model.Company
	db := gormdemo.GetDB()

	result := db.Model(&model.Company{}).Find(&codes)
	if result.Error != nil {
		return result.Error
	}

	var users []*model.User
	if err := db.Model(&model.User{}).Find(&users).Error; err != nil {
		return err
	}
	if len(users) == 0 {
		return nil
	}
	if len(codes) == 0 {
		return errors.New("cannot assign users: company table is empty")
	}

	for k, user := range users {
		user.CompanyID = codes[k%len(codes)].ID
		if err := db.Model(user).Update("company_code", user.CompanyID).Error; err != nil {
			return err
		}
	}
	return nil
}
