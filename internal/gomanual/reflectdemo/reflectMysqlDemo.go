package reflectdemo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type Device struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Port    int    `json:"port" db:"port"`
	// Auth
}

type Auth struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type PrintInfo interface {
	GetName() string
	GetAddress() string
	GetPort() int
	SetName(string)
	SetAddress(string)
	SetPort(int)
}

func (d *Device) GetName() string {
	return d.Name
}

func (d *Device) GetAddress() string {
	return d.Address
}

func (d *Device) GetPort() int {
	return d.Port
}

func (d *Device) SetName(name string) {
	d.Name = name
}

func (d *Device) SetAddress(address string) {
	d.Address = address
}

func (d *Device) SetPort(port int) {
	d.Port = port
}

// 通过反射的方式拼接Mysql字符串
func GenerateMysqlString(table interface{}) (string, []interface{}, error) {
	t := reflect.TypeOf(table)
	// 接口实际类型
	// fmt.Println(t)
	// 接口所属类型
	k := t.Kind()

	if k != reflect.Ptr {
		return "", nil, fmt.Errorf("reflect type must pointer.")
	}

	tableType := t.Elem()
	tableValue := reflect.ValueOf(table).Elem()

	if tableType.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("Only support struct type.")
	}

	var (
		sqlStr    bytes.Buffer
		sqlFields []string
		sqlParams []string
		sqlValues []interface{}
	)

	sqlStr.WriteString("INSERT INTO " + strings.ToLower(tableType.Name()) + " (")
	for i := 0; i < tableValue.NumField(); i++ {
		field := tableType.Field(i)
		value := tableValue.Field(i)

		// 指针可通过反射设置结构体的属性的值
		// if field.Type.Kind() == reflect.String && value.IsValid() && value.CanSet() {
		// 	value.SetString(fmt.Sprintf("Table %s", value))
		// }

		// 打印字段名称、字段类型、字段Tag、字段实际值 fmt.Printf("Field Name: %s, Field Type: %s, Field Tag: %s, Field Value: %v \n", field.Name, field.Type, field.Tag, value)
		sqlFields = append(sqlFields, fmt.Sprintf("`%s`", field.Tag.Get("db")))
		sqlParams = append(sqlParams, "?")
		sqlValues = append(sqlValues, value.Interface())
	}

	sqlStr.WriteString(fmt.Sprintf("%s ) VALUES (%s)", strings.Join(sqlFields, ","), strings.Join(sqlParams, ",")))

	return sqlStr.String(), sqlValues, nil
}
