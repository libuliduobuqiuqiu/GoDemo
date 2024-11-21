package storage

import (
	"godemo/internal/gostorage/gormdemo/dao"
	dao2 "godemo/internal/gostorage/gormgendemo/dao"
	"testing"
)

func TestGormFindUser1(t *testing.T) {
	if err := dao.ListUsersByTableName(); err != nil {
		t.Fatal(err)
	}
}

func TestGormFindUser2(t *testing.T) {
	if err := dao.ListUsersByNotTableName(); err != nil {
		t.Fatal(err)
	}
}

func TestGormGenFind(t *testing.T) {
	dao2.ListUsers()
}
