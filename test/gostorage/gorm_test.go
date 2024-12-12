package gostorage

import (
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormdemo/dao"
	"godemo/internal/gostorage/gormdemo/model"
	dao2 "godemo/internal/gostorage/gormgendemo/dao"
	"godemo/internal/gostorage/sqlxdemo"
	"testing"
)

func TestUseJoin(t *testing.T) {
	err := dao.UseGormJoin()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPreload(t *testing.T) {
	err := dao.UseGormPreload()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUserCompany(t *testing.T) {
	err := dao.UpdateUserCompany()
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertCompany(t *testing.T) {
	err := dao.InnsertCompanyRows()
	if err != nil {
		t.Fatal(err)
	}
}

func TestShowVariables(t *testing.T) {
	result, err := dao.ShowServiceVariables()
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range result {
		t.Log(k, v)
	}
}

func TestUseDryRun(t *testing.T) {
	dao.UseGromDryRun()
}

func TestInsertUser(t *testing.T) {
	if err := dao.InsertUser(); err != nil {
		t.Fatal(err)
	}
}

func TestGetUsers(t *testing.T) {
	err := dao.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransaction(t *testing.T) {
	if err := dao.InsertUserByTrans(); err != nil {
		t.Fatal(err)
	}
}

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

func TestExistDB(t *testing.T) {
	db, err := sqlxdemo.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	gormDB, err := gormdemo.InitDBByExistDB(db)
	if err != nil {
		t.Fatal(err)
	}

	var userList []*model.User
	result := gormDB.Table("users").Limit(10).Find(&userList)
	if result.Error != nil {
		t.Fatal(result.Error)
	}

	t.Log(result.RowsAffected)
	for _, v := range userList {
		t.Log(v)
	}
}
