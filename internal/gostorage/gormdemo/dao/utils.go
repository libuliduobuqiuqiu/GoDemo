package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"log"
	"strings"
)

// FilterTableField 筛选表中符合条件类型的字段
func FilterTableField(tableName []string) {
	type TableStruct struct {
		Field string `gorm:"column:Field"`
		Type  string `gorm:"column:Type"`
	}

	db := gormdemo.GetDB()
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
func ShowServiceVariables() (mysqlConfig map[string]string, err error) {
	type variable struct {
		VariableName string `gorm:"column:Variable_name"`
		Value        string `gorm:"column:value"`
	}

	var variableList []variable
	db := gormdemo.GetDB()
	result := db.Raw("show variables").Scan(&variableList)

	if result.Error != nil {
		return nil, result.Error
	}

	mysqlConfig = make(map[string]string)
	for _, v := range variableList {
		mysqlConfig[v.VariableName] = v.Value
	}
	return
}
