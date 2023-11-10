package gormDemo

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/pborman/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Company struct {
	ID   string `gorm:"column:id" faker:"-"`
	Code string `gorm:"column:code" faker:"word"`
	Name string `gorm:"column:name" faker:"word"`
}

type TableInfo struct {
	Schema string `gorm:"column:TABLE_SCHEMA" faker:"-"`
	Name   string `gorm:"column:TABLE_NAME" faker:"-"`
}

type TableStruct struct {
	Field string `gorm:"column:Field"`
	Type  string `gorm:"column:Type"`
}

func (t *TableInfo) TableName() string {
	return "tables"
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

// GenerateSQL 生成修改表名SQL
func GenerateSQL(db *gorm.DB) {
	var tables []TableInfo

	result := db.Where(&TableInfo{Schema: "adops"}).Find(&tables)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)

	for _, t := range tables {
		newTable := fmt.Sprintf("instance_ad_%s", t.Name)
		sql := fmt.Sprintf("RENAME TABLE `%s`.`%s` TO `%s`.`%s`;", "adops", t.Name, "adops", newTable)
		// sql := fmt.Sprintf("RENAME TABLE `%s`.`%s` TO `%s`.`%s`;", "adops", newTable, "adops", t.Name)
		fmt.Println(sql)
	}

}

// FilterADTable 筛选有关Adops的表
func FilterADTable(db *gorm.DB) (tableNames []string) {
	var tables []TableInfo

	result := db.Where(&TableInfo{Schema: "cmdb"}).Find(&tables)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	for _, t := range tables {
		if strings.Contains(t.Name, "instance_ad_") {
			tableNames = append(tableNames, t.Name)
		}

	}
	return
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

func GeneratePoolMember(db *gorm.DB) {

	type VirtualServer struct {
		ID         string `gorm:"id"`
		SynGroupID string `gorm:"syn_group_id""`
	}

	var vsList []VirtualServer
	result := db.Raw("select id, syn_group_id from dns_virtual_server where syn_group_id = '1695982405669892096'").Scan(&vsList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	type PoolMember struct {
	}

}

func CompareMysqlConfig() {
	db, err := InitDB("company")
	if err != nil {
		fmt.Println(err)
		return
	}

	mysqlConfig1 := ShowServiceVariables(db)

	db2, err := InitDB("company2")
	if err != nil {
		fmt.Println(err)
		return
	}

	mysqlConfig2 := ShowServiceVariables(db2)

	if mysqlConfig1 != nil && mysqlConfig2 != nil {
		for k, v1 := range mysqlConfig1 {
			if v2, ok := mysqlConfig2[k]; ok {
				if v1 != v2 {
					fmt.Printf("配置参数的%s不一致，75Mysql：%s，57Mysql：%s\n", k, v1, v2)
				}
				continue
			}
			fmt.Println("配置参数不存在：", k)

		}
	}

}
