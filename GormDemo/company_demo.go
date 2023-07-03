package gormDemo

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/pborman/uuid"
	"gorm.io/gorm"
	"log"
)

type Company struct {
	ID   string `gorm:"column:id" faker:"-"`
	Code string `gorm:"column:code" faker:"word"`
	Name string `gorm:"column:name" faker:"word"`
}

func (c *Company) TableName() string {
	return "company"
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}

// insertRows 批量插入
func insertCompanyRows(db *gorm.DB) {
	var companys []*Company
	for i := 0; i < 10; i++ {
		tmpCompany := Company{}
		err := faker.FakeData(&tmpCompany)
		if err != nil {
			log.Fatal(err.Error())
		}
		companys = append(companys, &tmpCompany)
	}

	result := db.Create(companys)

	if result.Error != nil {
		log.Fatalf(result.Error.Error())
	}

	fmt.Println("RowsAffected: ", result.RowsAffected)

	for _, m := range companys {
		fmt.Printf("%+v\n", m)
	}
}

func SelectCompanyRows(db *gorm.DB) {
	var codes []Company

	result := db.Model(&Company{}).Find(&codes)
	fmt.Println(codes)
	fmt.Println(result.Error, result.RowsAffected)

	var users []*User
	db.Model(&User{}).Find(&users)

	for k, user := range users {
		user.CompanyID = codes[k%len(codes)].Code

	}

	for _, tmpUser := range users {
		fmt.Println(tmpUser)
		result = db.Model(&User{}).Where("id = ?", tmpUser.ID).Updates(tmpUser)
		fmt.Println(result.Error, result.RowsAffected)
	}

}
