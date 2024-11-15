package gormdemo

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-faker/faker/v4"
	"godemo/internal/gostorage/gormdemo/model"
	"gorm.io/gorm"
)

type TableStruct struct {
	Field string `gorm:"column:Field"`
	Type  string `gorm:"column:Type"`
}

// insertRows 批量插入
func insertCompanyRows(db *gorm.DB) {
	var companys []*model.Company
	for i := 0; i < 10; i++ {
		tmpCompany := model.Company{}
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

// SelectCompanyRows 查询单表后更新
func SelectCompanyRows(db *gorm.DB) {
	var codes []*model.Company

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

// FilterTableField 筛选表中符合条件类型的字段
func FilterTableField(db *gorm.DB, tableName []string) {
	for _, name := range tableName {
		sql := fmt.Sprintf("DESCRIBE %s", name)
		var tableFieldList []TableStruct
		result := db.Raw(sql).Scan(&tableFieldList)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		for _, field := range tableFieldList {
			if !strings.Contains(field.Type, "char") && !strings.Contains(field.Type, "int") {
				fmt.Println(name, field.Field, field.Type)
			}
		}

	}
}

// ShowServiceVariables 查看当前数据库中的环境变量
func ShowServiceVariables(db *gorm.DB) (mysqlConfig map[string]string) {

	type variable struct {
		VariableName string `gorm:"column:Variable_name"`
		Value        string `gorm:"column:value"`
	}

	var variableList []variable
	result := db.Raw("show variables").Scan(&variableList)

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	mysqlConfig = make(map[string]string)
	for _, v := range variableList {
		mysqlConfig[v.VariableName] = v.Value
	}
	return
}
