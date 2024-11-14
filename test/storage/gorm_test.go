package storage

import (
	"godemo/internal/gostorage/gormdemo/dao"
	dao2 "godemo/internal/gostorage/gormgendemo/dao"
	"testing"
)

func TestGormFind(t *testing.T) {

	dao.ListUsers()

}

func TestGormGenFind(t *testing.T) {
	dao2.ListUsers()
}
