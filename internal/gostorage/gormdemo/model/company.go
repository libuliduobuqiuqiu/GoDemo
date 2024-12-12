package model

import (
	"github.com/pborman/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID    string `gorm:"column:id" faker:"-"`
	Code  string `gorm:"column:code" faker:"word"`
	Name  string `gorm:"column:name" faker:"word"`
	Users []User `gorm:"foreignKey: CompanyID"`
}

func (c *Company) TableName() string {
	return "company"
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}
